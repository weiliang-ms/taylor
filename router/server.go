package router

import (
	"github.com/gin-gonic/gin"
	"taylor/authentic"
	"taylor/nginx"
	"taylor/utils"
	"taylor/waf"
)

func Router() {

	r := gin.New()
	// 全局日志中间件
	r.Use(LoggerMiddleware())
	//r.Use(authentic.Filter())
	r.Static("/static", "static/")
	// nginx控制器路由组
	nginx.AddRoute(r)
	// waf控制器路由
	waf.AddRoute(r)

	instance := r.Group("/v1/process/instance")
	{
		//instance.GET("/info", nginx.InstanceInfo)
		instance.GET("/start", nginx.Start)
		instance.GET("/stop", nginx.Stop)
		instance.GET("/reload", nginx.Reload)
	}

	auth := r.Group("/v1/authorization")
	{
		auth.POST("/login", authentic.Login)
	}

	err := r.Run(":8099")
	if err != nil {
		utils.Logger.Errorln(err)
	}
}
