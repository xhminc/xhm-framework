package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/xhminc/xhm-framework/config"
	"strings"
	"time"
)

func BuildJwt(key interface{}, claims jwt.MapClaims) (string, error) {

	var method jwt.SigningMethod
	globalConfig = config.GetGlobalConfig()
	method = GetSignMethod(globalConfig.Application.Jwt.Method)
	claims["exp"] = time.Now().Add(*globalConfig.Application.Jwt.Timeout).Unix()

	token := jwt.NewWithClaims(method, claims)
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func GetSignMethod(method string) jwt.SigningMethod {

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
