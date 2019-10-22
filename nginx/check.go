package nginx

import (
	"taylor/utils"
	"os"
	"regexp"
	"strconv"
)

// 检测端口合法性
func checkPort(port int) (msg string, result bool) {

	// 检测端口取值范围
	if port < 1 || port > 65535 {
		return utils.PortOutOfRange, false
	}

	// 检测端口是否被占用
	if _, code := utils.Shell(utils.CheckPortDirect + strconv.Itoa(port)); code == 0 {
		return utils.PortInUse, false
	}

	return utils.PortCheckSuccessful, true
}

// 检测配置文件名称
func checkConfigName(name string, filePath string) (msg string, result bool) {

	reg := regexp.MustCompile("^[a-z0-9A-Z\\-\\_]+$")
	if !reg.MatchString(name) {
		return utils.ConfigNameIllegal, false
	}
	fileInfo, _ := os.Stat(filePath)
	if fileInfo != nil {
		return utils.ConfigNameExist, false
	}
	return utils.ConfigNameCheckSuccessful, true
}
