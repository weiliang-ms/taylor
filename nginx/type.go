package nginx

import (
	"taylor/utils"
	"os"
	"strings"
)

// 根据路径匹配配置文件类型
func parseConfigType(path string) string {
	fileInfo, err := os.Stat(path)
	if err != nil {
		utils.Logger.Errorln(err)
	}
	if !fileInfo.IsDir() && fileInfo.Name() == "nginx.conf" {
		return utils.CoreType
	}

	if strings.HasSuffix(fileInfo.Name(), utils.BackupSuffix) {
		return utils.DisableType
	}

	dir := strings.Trim(path, fileInfo.Name())
	switch dir {
	case CoreConfigPath():
		return utils.CoreType
	case HttpConfigDir():
		return utils.HttpType
	case StreamConfigDir():
		return utils.TcpType
	case WafConfigDir():
		return utils.WafType
	case WafWhiteConfigDir():
		return utils.WafType
	case WafBlackConfigDir():
		return utils.WafType
	default:
		return utils.OtherType
	}

}
