package nginx

import (
	"bufio"
	"io"
	"taylor/utils"
	"os"
	"strings"
)

type ConfigInstance struct {
	name    string
	form    string // 配置类型
	path    string
	port    string
	content []byte
}

// 过滤查询
func searchConfigs(instance SearchInstance, page int, limit int) (count int, data []map[string]interface{}) {
	var afterFilterConfigs []string
	for _, f := range ConfigList() {
		fileInfo, _ := os.Stat(f)
		port := parseListenPort(parseConfigType(f), f)
		portFilter := strings.Contains(port, instance.KeyWord)
		typeFilter := strings.Contains(parseConfigType(f), instance.KeyWord)
		nameFilter := strings.Contains(fileInfo.Name(), instance.KeyWord)
		if portFilter || typeFilter || nameFilter {
			afterFilterConfigs = append(afterFilterConfigs, f)
		}
	}
	return len(afterFilterConfigs), packageConfigs(utils.Paging(page, limit, afterFilterConfigs))
}

// 组装nginx 配置实例
func packageConfigs(configs []string) (data []map[string]interface{}) {
	var slice []map[string]interface{}
	id := 0
	for _, f := range configs {
		id++
		fileInfo, _ := os.Stat(f)
		modifyTime := fileInfo.ModTime().Format("2006-01-02 15:04:05")
		var config map[string]interface{}
		config = make(map[string]interface{})
		config["id"] = id
		config["name"] = fileInfo.Name()
		config["form"] = parseConfigType(f)
		config["path"] = f
		config["port"] = parseListenPort(parseConfigType(f), f)
		config["content"] = string(readFileContent(f))
		config["modifyTime"] = modifyTime
		slice = append(slice, config)
	}

	return slice
}

// 解析监听端口
func parseListenPort(configType string, path string) string {
	var port string
	if configType == utils.WafType || configType == utils.OtherType || strings.HasSuffix(path, ".bak") {
		return "/"
	}
	var listenPort string
	f, err := os.Open(path)
	if err != nil {
		utils.Logger.Errorln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			if listenPort == "" {
				listenPort = "/"
			}
			break
		}
		if strings.Contains(line, "listen") {
			port = utils.DeleteSubString(line, "listen", ";", "\n", "\t", "\r")
			listenPort = listenPort + port + " "
		}

	}
	return listenPort
}

// 解析nginx worker数量
func getNginxWorkerNum() (workerNum string) {
	workerNum = utils.ReadFileKeyWord(CoreConfigPath(), "worker_processes")
	switch workerNum {
	case "default":
		return workerNum
	default:
		return utils.DeleteSubString(workerNum, "worker_processes", ";")
	}
}
