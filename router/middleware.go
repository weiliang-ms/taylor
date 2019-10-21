package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"taylor/utils"
	"time"
)

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
