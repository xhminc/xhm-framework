package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"github.com/xhminc/xhm-framework/xhm"
	"os"
	"time"
)

var (
	timestampFormat = "2006-01-02 15:04:05.000"
)

func LogToFile() gin.HandlerFunc {

	var log = logrus.New()

	if xhm.GlobalConfig.Application.Profile == "dev" || xhm.GlobalConfig.Application.Profile == "test" {
		log.SetFormatter(&prefixed.TextFormatter{
			ForceColors:      true,
			ForceFormatting:  true,
			FullTimestamp:    true,
			DisableUppercase: true,
			TimestampFormat:  timestampFormat,
		})
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)
	} else {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timestampFormat,
		})
		log.SetLevel(logrus.InfoLevel)
		log.SetOutput(os.Stdout)
	}

	log.Info("LoggerToFile component inited")

	return func(c *gin.Context) {

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		log.Info(latencyTime)
		log.Info(reqMethod)
		log.Info(reqUri)
		log.Info(statusCode)
		log.Info(clientIP)
	}
}
