package security

import (
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/auth"
	"github.com/xhminc/xhm-framework/component/base"
	"github.com/xhminc/xhm-framework/component/common"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"net/http"
)

const (
	XHM_TOKEN = "XHM-Token"
)

var (
	sessionTimeout = base.Result{
		Code:    90000,
		Message: "common.sessionTimeout",
	}
)

func SessionHandler() gin.HandlerFunc {

	log = logger.GetLogger()
	globalConfig = config.GetGlobalConfig()
	log.Info("Loading \"security.SessionHandler\" component success")

	return func(c *gin.Context) {

		requestUri := c.FullPath()
		for _, url := range globalConfig.Application.Session.IgnoreUrls {
			if requestUri == url {
				c.Next()
				return
			}
		}

		xhmToken := c.GetHeader(XHM_TOKEN)
		if len(xhmToken) == 0 {
			c.AbortWithStatusJSON(http.StatusOK, sessionTimeout)
			return
		}

		if auth.IsTokenValid(xhmToken) {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusOK, sessionTimeout)
		}

	}
}
