package waf

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	wafRouter := r.Group("/v1/process/waf")
	//wafRouter.Use(nginx.Filter())
	//{
	wafRouter.POST("/rules", WhiteHost)
	wafRouter.POST("/edict", Modify)
	wafRouter.POST("/add", AddRule)
	wafRouter.POST("/delete", DeleteRule)
	//}
}
