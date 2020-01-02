package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/test"
	"testing"
)

func TestLogToFile(t *testing.T) {

	r := gin.New()
	r.Use(LogToFile())

	r.GET("/", func(c *gin.Context) {

	})

	test.PerformRequest(r, "GET", "/")
}
