package security

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

var (
	log          *zap.Logger
	globalConfig *config.YAMLConfig
)

func CorsRequestHandler() gin.HandlerFunc {

	log = logger.GetLogger()
	globalConfig = config.GetGlobalConfig()
	log.Info("Loading \"security.CorsRequestHandler\" component success")

	return func(c *gin.Context) {

		method := c.Request.Method

		if len(globalConfig.Application.Cors.AccessControlAllowOrigin) > 0 {
			c.Header("Access-Control-Allow-Origin",
				strings.Join(globalConfig.Application.Cors.AccessControlAllowOrigin, ","))
		}

		if len(globalConfig.Application.Cors.AccessControlAllowMethods) > 0 {
			c.Header("Access-Control-Allow-Methods",
				strings.Join(globalConfig.Application.Cors.AccessControlAllowMethods, ","))
		}

		if len(globalConfig.Application.Cors.AccessControlAllowHeaders) > 0 {
			c.Header("Access-Control-Allow-Headers",
				strings.Join(globalConfig.Application.Cors.AccessControlAllowHeaders, ","))
		}

		if len(globalConfig.Application.Cors.AccessControlExposeHeaders) > 0 {
			c.Header("Access-Control-Expose-Headers",
				strings.Join(globalConfig.Application.Cors.AccessControlExposeHeaders, ","))
		}

		if globalConfig.Application.Cors.AccessControlAllowCredentials != nil {
			c.Header("Access-Control-Allow-Credentials",
				strconv.FormatBool(*globalConfig.Application.Cors.AccessControlAllowCredentials))
		}

		if globalConfig.Application.Cors.AccessControlMaxAge != nil {
			c.Header("Access-Control-Max-Age",
				strconv.FormatFloat(globalConfig.Application.Cors.AccessControlMaxAge.Seconds(), 'E', -1, 64))
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}

	}
}
