package nginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taylor/utils"
	"strings"
)

// nginx配置调整中间件
func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		utils.Logger.Info("api接口：", c.Request.RequestURI)
		utils.Logger.Info("上个请求状态码：", c.Writer.Status())
		if !strings.Contains(c.Request.RequestURI, "select") && c.Writer.Status() == http.StatusOK {
			Reload(c)
		}
	}
}
