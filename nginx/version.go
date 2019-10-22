package nginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"taylor/utils"
)

const systemName = "taylor"

type VersionInfo struct {
	Name string
	Version string
	SecurityModule string
	OtherModule string
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
	return info
}

func readVersion() (v string)  {
	// 从环境变量读取
	v = os.Getenv("NGINX_VERSION")
	if v != ""{
		return v
	}

	// 从文件读取
	v = utils.ReadContent()
}
