package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xhminc/xhm-framework/component/common"
	"github.com/xhminc/xhm-framework/component/logger"
	"github.com/xhminc/xhm-framework/config"
	"net/http"
)

const (
	XHM_TOKEN = "XHM-Token"
)

var (
	sessionTimeout = common.Result{
		Code:    90000,
		Message: "global.session.timeout",
	}
)

func SessionHandler() gin.HandlerFunc {

	log = logger.GetLogger()
	globalConfig = config.GetGlobalConfig()
	log.Info("Loading \"security.SessionHandler\" component success")

	return func(c *gin.Context) {

		requestUri := c.Request.RequestURI
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

		token, err := jwt.Parse(xhmToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(globalConfig.Application.Jwt.Key), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, sessionTimeout)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusOK, sessionTimeout)
		}
	}
}
