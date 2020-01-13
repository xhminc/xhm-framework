package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/xhm"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func RequestLogger() gin.HandlerFunc {

	log.Info("Loading \"logger.RequestLogger\" component success")

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

		log.Info(
			reqMethod+" "+reqUri,
			zap.Int("status", statusCode),
			zap.Duration("cost", latencyTime),
			zap.String("ip", clientIP),
			zap.Any("params", c.Params),
			zap.Any("query", c.Request.URL.Query()),
			zap.Any("headers", getRequestHeaders(c)),
		)
	}
}

func getRequestHeaders(c *gin.Context) http.Header {
	var header http.Header
	if xhm.IsProduct() {
		header = c.Request.Header
	}
	return header
}
