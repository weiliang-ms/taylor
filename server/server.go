package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"taylor/authentic"
	"taylor/nginx"
	"taylor/utils"
	"taylor/waf"
	"time"
)

func Start() {

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

// 日志记录中间件
func LoggerMiddleware() gin.HandlerFunc {

	//实例化
	logger := utils.Logger

	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 日志格式
		logger.Infof("%d | %v | %s | %s | %s",
			c.Writer.Status(),
			c.ClientIP(),
			latencyTime,
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}
