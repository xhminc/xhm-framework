package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/component/test"
	"testing"
	"time"
)

var (
	timestampFormat = "20060102150405000"
)

func TestLogToFile(t *testing.T) {

	r := gin.New()
	r.Use(logger.RequestLogger())

	r.GET("/:id", func(c *gin.Context) {
	})

	ts := time.Now().Format(timestampFormat)
	test.PerformRequest(r, "GET", "/887910?t="+ts+"&name=service&name=xhm")
}
