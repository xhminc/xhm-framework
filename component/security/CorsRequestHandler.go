package security

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"net/http"
	"strings"
)

var (
	log          = logger.GetLogger()
	globalConfig = config.GetGlobalConfig()
)

func CorsRequestHandler() gin.HandlerFunc {

	log.Info("Loading \"security.CorsRequestHandler\" component success")

	return func(c *gin.Context) {

		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", strings.Join(globalConfig.Application.Cors.Hosts, ","))
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Token")
		c.Header("Access-Control-Expose-Headers",
			"Content-Type, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		c.Header("Content-Type", "application/json")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
