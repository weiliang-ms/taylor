package nginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"taylor/utils"
)

const systemName = "taylor"
const securityModules  = "waf | naxsi | badbots"
const coreModules  = "stream | ssl | more_header"
const features  = "日志切割 | 安全防护 | 监控管理 | web"

type VersionInfo struct {
	Name string
	Version string
	SecurityModule string
	CoreModule string
	Feature string
}

// nginx版本信息
func Version(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data":  info(),
	})
}
func info() map[string]string  {
	info := make(map[string]string, 5)
	info["name"] = systemName
	info["version"] = readVersion()
	info["security"] = securityModules
	info["core"] = coreModules
	info["feature"] = features
	return info
}

func readVersion() (v string)  {
	// 从环境变量读取
	if v = os.Getenv("NGINX_VERSION"); v != ""{
		utils.Logger.Infof("nginx版本为：", v)
		return v
	}

	// 从文件读取
	if v = utils.ReadContent(VersionConfigPath()); v != "" {
		utils.Logger.Infof("nginx版本为：", v)
		return v
	}

	// TODO:从二进制解析
	return "1.16.0"
}
