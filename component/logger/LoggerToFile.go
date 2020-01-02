package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerToFile() gin.HandlerFunc {

	fmt.Println("Logger interceptor init...")

	return func(c *gin.Context) {

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		log.Info(latencyTime)
	}
}
