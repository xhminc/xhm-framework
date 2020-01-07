package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/component/test"
	"github.com/xhminc/xhm-framework/xhm"
	"testing"
	"time"
)

var (
	timestampFormat = "20060102150405000"
)

func TestLogToFile(t *testing.T) {

	xhm.Bootstrap()

	r := gin.New()
	r.Use(logger.RequestLogger())

	r.GET("/", func(c *gin.Context) {
	})

	ts := time.Now().Format(timestampFormat)
	test.PerformRequest(r, "GET", "/?t="+ts)
}
