package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func RequestLogger() gin.HandlerFunc {

	log.Info("Loading \"logger.RequestLogger\" component success")

	return func(c *gin.Context) {

		startTime := time.Now()
		reqMethod := c.Request.Method
		reqUri := c.FullPath()
		clientIP := c.ClientIP()

		log.Info(
			reqMethod+" [request] "+reqUri,
			zap.String("ip", clientIP),
			zap.Any("params", c.Params),
			zap.Any("query", c.Request.URL.Query()),
			zap.Any("headers", getRequestHeaders(c)),
		)

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		statusCode := c.Writer.Status()

		log.Info(
			reqMethod+" [response] "+reqUri,
			zap.Int("status", statusCode),
			zap.Duration("cost", latencyTime),
		)
	}
}

func getRequestHeaders(c *gin.Context) http.Header {
	var header http.Header
	if globalConfig.IsProduct() {
		header = c.Request.Header
	}
	return header
}
