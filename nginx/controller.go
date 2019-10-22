package nginx

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"taylor/utils"
)

// 查询
func all(page int, limit int) (count int, data []map[string]interface{}) {
	return getConfigs(page, limit)
}

// 过滤查询
func Select(c *gin.Context) {
	var instance SearchInstance
	var data []map[string]interface{}
	var count int
	c.ShouldBind(&instance)
	page, _ := strconv.Atoi(instance.Page)
	limit, _ := strconv.Atoi(instance.Limit)
	if instance.KeyWord == "all" {
		count, data = all(page, limit)
	} else {
		count, data = searchConfigs(instance, page, limit)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
	})
}

// 重命名文件
func Rename(c *gin.Context) {
	var instance SearchInstance
	c.BindJSON(&instance)
	err := os.Rename(instance.Path, utils.DirAppendSlash(utils.CurrentDir(instance.Path))+instance.Name)
	if err != nil {
		utils.Logger.Errorln(err.Error())
		response(c, InternalError, err.Error())
	}
}

// 调整文件名称，间接管理配置
func Change(c *gin.Context) {

	// 获取body数据
	var instance SearchInstance
	c.BindJSON(&instance)
	// 修改后配置文件属性，回执用于重载
	name := changeName(instance.Name)
	path := utils.DirAppendSlash(utils.CurrentDir(instance.Path))
	if instance.Name == "nginx.conf" || instance.Name == "init.lua" || instance.Name == "access.lua" {
		response(c, Forbidden, utils.CoreConfigNotAllowModify)
	} else {
		utils.Logger.Info("更改配置文件启用状态,修改前为：" + instance.Name)
		err := os.Rename(instance.Path, path+name)
		if err != nil {
			response(c, InternalError, err.Error())
		} else {
			utils.Logger.Info("更改配置文件启用状态,修改后为：" + name)
		}
	}
}

// 删除文件
func Delete(c *gin.Context) {
	var instance SearchInstance
	c.BindJSON(&instance)
	err := deleteConfigs(instance)
	if err != nil {
		response(c, InternalError, err.Error())
	}
}

// 编辑配置文件
func Update(c *gin.Context) {

	var instance ConfigDetails
	c.BindJSON(&instance)
	f, err := os.OpenFile(instance.Path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		utils.Logger.Errorln(err.Error())
		response(c, InternalError, err.Error())
	}

	if _, err = f.Write([]byte(instance.Content)); err != nil {
		utils.Logger.Errorln(err.Error())
		response(c, InternalError, err.Error())
	}
}

// 自定义配置文件
func Custom(c *gin.Context) {
	// 解析form表单
	var instance ConfigDetails
	c.BindJSON(&instance)

	// 检测端口合法性
	msg, port := checkPort(instance.Port)
	if !port {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  msg,
		})
		return
	}

	// 检测名称合法性
	// 组装配置文件path
	var filePath string
	switch instance.Protocol {
	case utils.HttpType:
		filePath = HttpConfigDir() + instance.Name + utils.HttpConfigSuffix
	case utils.TcpType:
		filePath = StreamConfigDir() + instance.Name + utils.TcpConfigSuffix
	}
	msg, name := checkConfigName(instance.Name, filePath)

	if !name {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  msg,
		})
		return
	} else {
		// 写文件并判断
		_, code := writeConfigContent(filePath, instance.Content)
		if code != 0 {
			delErr := os.Remove(filePath)
			if delErr != nil {
				utils.Logger.Errorln(delErr.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 0,
					"msg":  delErr.Error(),
				})
				return
			}
		}
	}
}

// 上传配置文件
func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	var path string
	if strings.HasSuffix(header.Filename, utils.HttpConfigSuffix) {
		path = utils.DirAppendSlash(HttpConfigDir()) + header.Filename
	} else if strings.HasSuffix(header.Filename, utils.TcpConfigSuffix) {
		path = utils.DirAppendSlash(StreamConfigDir()) + header.Filename
	}

	fileInfo, _ := os.Stat(path)
	if fileInfo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  utils.FileExists,
		})
	} else {
		targetFile, createErr := os.Create(path)
		if createErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 1,
				"msg":  createErr.Error(),
			})
		}
		_, err = io.Copy(targetFile, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 1,
				"msg":  createErr.Error(),
			})
		}
		defer targetFile.Close()
	}
}
