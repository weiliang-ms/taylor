package nginx

import (
	"github.com/gin-gonic/gin"
)

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
}
