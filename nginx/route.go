package nginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"taylor/utils"
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

func AddRoute(r *gin.Engine) {
	conf := r.Group("/v1/process/conf")
	conf.Use(Filter())
	{
		//conf.GET("/select", SelectAll)
		conf.POST("/select", Select)
		conf.POST("/update", Update)
		conf.POST("/delete", Delete)
		conf.POST("/rename", Rename)
		conf.POST("/change", Change)
		conf.POST("/upload", Upload)
		conf.POST("/custom", Custom)
	}

	instance := r.Group("/v1/process/instance")
	{
		instance.GET("/version", Version)
	}
}
