package waf

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"taylor/nginx"
	"taylor/utils"
	"os"
	"runtime"
	"strings"
)

const (
	BaseConfigDirVariable = "NginxWafConfigDir"
	UnixDefaultWafDir     = nginx.UnixDefaultBaseDir + "conf/waf/"
	WindowsDefaultWafDir  = nginx.WindowsDefaultBaseDir + "conf\\waf\\"
)

type ProcessInstance struct {
	Type    string `form:"type"`
	Rule    string `form:"rule"`
	Page    int    `form:"page"`
	Content string `form:"content"`
	Limit   int    `form:"limit"`
}

type ModifyInstance struct {
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
	Path     string `json:"path"`
}

// 获取waf配置文件路径
func BaseConfigDir() string {
	if os.Getenv(BaseConfigDirVariable) == "" {
		switch runtime.GOOS {
		case "windows":
			return WindowsDefaultWafDir
		default:
			return UnixDefaultWafDir
		}
	}
	return utils.DirAppendSlash(os.Getenv(BaseConfigDirVariable))
}

func WhiteHost(c *gin.Context) {
	var instance ProcessInstance
	c.Bind(&instance)
	path := utils.AppendPathWithSlash(BaseConfigDir(), instance.Type, instance.Rule)
	// TODO:后期添加分页
	wafRules := packageWafRules(instance.Page, instance.Limit, path, instance.Rule)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"error": nil,
		"count": len(rules(path)),
		"data":  wafRules,
	})
}

// 组装waf规则
func packageWafRules(page int, limit int, path string, configRule string) (data []map[string]interface{}) {
	var slice []map[string]interface{}
	id := 0
	fileInfo, _ := os.Stat(path)
	modifyTime := fileInfo.ModTime().Format("2006-01-02 15:04:05")
	status := ruleStatus(configRule, fileInfo.Name())
	for _, rule := range utils.Paging(page, limit, rules(path)) {
		id++
		var config map[string]interface{}
		config = make(map[string]interface{})
		config["id"] = id
		config["modifyTime"] = modifyTime
		config["status"] = status
		config[configRule] = rule
		config["path"] = path
		slice = append(slice, config)
	}

	return slice
}

// 读取规则
func rules(path string) (rules []string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		utils.Logger.Errorln(err.Error())
	}

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		utils.Logger.Info("读取内容：", "ss"+line+"ss")
		if line != utils.Linefeed() && line != "" {
			rules = append(rules, line)
		}
		if err != nil || err == io.EOF {
			utils.Logger.Errorln(err.Error())
			break
		}
	}

	return
}

func Modify(c *gin.Context) {
	var instance ModifyInstance
	c.BindJSON(&instance)
	//TODO: 后续添加正则校验新value合法性
	utils.ModifyContent(instance.Path, instance.OldValue, instance.NewValue)
}

// 添加waf具体规则
func AddRule(c *gin.Context) {
	var re error
	var msg string
	var instance ProcessInstance
	c.BindJSON(&instance)
	path := utils.AppendPathWithSlash(BaseConfigDir(), instance.Type, instance.Rule)

	// 检测rule规则格式是否正确
	if instance.Rule == "method" {
		instance.Content = strings.ToUpper(instance.Content)
	}

	if checkRule(instance.Rule, instance.Content) {
		re, msg = utils.AddLine(path, instance.Content)
	} else {
		re, msg = utils.IllegalFormatError, utils.IllegalFormatError.Error()
	}

	c.JSON(utils.ConvertHttpStatus(re), gin.H{
		"code":  0,
		"error": nil,
		"msg":   msg,
	})
}

// TODO:后续修改为操作文件方式，而非调用shell方式
func DeleteRule(c *gin.Context) {
	var instance ProcessInstance
	c.BindJSON(&instance)
	path := utils.AppendPathWithSlash(BaseConfigDir(), instance.Type, instance.Rule)
	msg, code := utils.Shell("sed -i \"/" + utils.DeleteSubString(instance.Content, "\n", "\r") + "/d\" " + path)
	utils.Logger.Info(msg, code)
}

// 检测规则格式是否正确
func checkRule(rule string, ruleContent string) bool {
	switch rule {
	case "host":
		return utils.CheckHostAddr(ruleContent)
	case "method":
		return utils.CheckMethod(ruleContent)
	default:
		return true
	}
}

// waf规则是否被激活
func ruleStatus(rule string, fileName string) string {
	if strings.HasSuffix(fileName, utils.BackupSuffix) {
		return utils.DisableType
	}
	// 解析总开关
	if Status("core") == utils.DisableType || Status(rule) == utils.DisableType {
		return utils.DisableType
	}
	return utils.ActiveStatus
}

// 解析waf状态
func Status(rule string) string {
	status := utils.ReadFileKeyWord(nginx.WafConfigDir()+"config.lua", ruleSwitchName(rule))
	switch utils.DeleteSubString(status, ruleSwitchName(rule), "\r", "\n") {
	case "true":
		return utils.ActiveStatus
	default:
		return utils.DisableType
	}
}

// 获取waf开关状态（waf/config.lua）
func ruleSwitchName(ruleType string) string {
	switch ruleType {
	case "core":
		return "open_waf="
	default:
		return "open_check_" + ruleType + "="
	}
}
