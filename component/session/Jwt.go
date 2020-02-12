package session

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xhminc/xhm-framework/config"
	"strings"
	"time"
)

type cookie struct {
	Name     string `json:"name"`
	Timeout  int64  `json:"timeout"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	HttpOnly bool   `json:"httpOnly"`
}

func BuildJwt(key string, payload interface{}) (string, error) {

	var method jwt.SigningMethod
	globalConfig = config.GetGlobalConfig()
	method = getSignMethod(globalConfig.Application.Session.Jwt.Method)

	mapClaims := jwt.MapClaims{
		"exp":     time.Now().Add(*globalConfig.Application.Session.Jwt.Timeout).Unix(),
		"payload": payload,
	}

	if globalConfig.Application.Session.Cookie.Enable {
		mapClaims["cookie"] = cookie{
			Name:     globalConfig.Application.Session.Cookie.Name,
			Timeout:  int64(globalConfig.Application.Session.Cookie.Timeout.Seconds()),
			Path:     globalConfig.Application.Session.Cookie.Path,
			Domain:   globalConfig.Application.Session.Cookie.Domain,
			HttpOnly: globalConfig.Application.Session.Cookie.HttpOnly,
		}
	}

	token := jwt.NewWithClaims(method, mapClaims)
	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func getSignMethod(method string) jwt.SigningMethod {

	var signingMethod jwt.SigningMethod

	switch strings.ToUpper(method) {
	case "ES256":
		signingMethod = jwt.SigningMethodES256
		break
	case "HS256":
		signingMethod = jwt.SigningMethodHS256
		break
	case "PS256":
		signingMethod = jwt.SigningMethodPS256
		break
	case "ES384":
		signingMethod = jwt.SigningMethodES384
		break
	case "HS384":
		signingMethod = jwt.SigningMethodHS384
		break
	case "PS384":
		signingMethod = jwt.SigningMethodPS384
		break
	case "ES512":
		signingMethod = jwt.SigningMethodES512
		break
	case "HS512":
		signingMethod = jwt.SigningMethodHS512
		break
	case "PS512":
		signingMethod = jwt.SigningMethodPS512
		break
	}

	return signingMethod
}
