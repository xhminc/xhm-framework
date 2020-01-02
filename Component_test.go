package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/component/test"
	"github.com/xhminc/xhm-framework/config"
	"github.com/xhminc/xhm-framework/xhm"
	"testing"
)

func TestLogToFile(t *testing.T) {

	CommonConfig := config.YAMLConfig{}
	xhm.Bootstrap(&CommonConfig)

	r := gin.New()
	r.Use(logger.LogToFile())

	r.GET("/", func(c *gin.Context) {
	})

	test.PerformRequest(r, "GET", "/")
}
