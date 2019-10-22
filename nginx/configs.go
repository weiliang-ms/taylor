package nginx

import (
	"taylor/errors"
	"taylor/utils"
	"os"
	"strings"
)

// 前台操作配置的相关信息（自定义配置、上传配置接口）
type ConfigDetails struct {
	Path       string `json:"path"`       // 配置文件路径
	Content    string `json:"content"`    // 配置文件内容
	Protocol   string `json:"protocol"`   // nginx代理的协议
	Naxsi      string `json:"naxsi"`      // 是否开启naxsi模块
	ExposePort string `json:"exposePort"` // 防火墙是否开放端口
	Port       int `json:"port"`       // 配置监听端口
	Name       string `json:"name"`       // 配置文件名称
}

type SearchInstance struct {
	Limit   string `form:"limit"`
	Page    string `form:"page"`
	KeyWord string `form:"keyWord"`
	Path    string `form:"path"`
	Name    string `form:"name"`
}

// 全量查询
func getConfigs(page int, limit int) (count int, data []map[string]interface{}) {
	return len(ConfigList()), packageConfigs(utils.Paging(page, limit, ConfigList()))
}

// 文件.bak后缀转换
func changeName(name string) (newName string) {
	if strings.Contains(name, utils.BackupSuffix) {
		newName = strings.Replace(name, utils.BackupSuffix, "", -1)
	} else {
		newName = name + utils.BackupSuffix
	}
	return
}

// 删除配置
func deleteConfigs(instance SearchInstance) error {
	fileInfo, err := os.Stat(instance.Path)
	if err != nil || fileInfo == nil {
		return errors.New("删除失败！")
	}
	return os.Remove(instance.Path)
}
