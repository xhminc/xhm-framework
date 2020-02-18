package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/xhminc/xhm-framework/config"
	"github.com/xhminc/xhm-framework/util/tripledes"
	"strings"
	"time"
)

func IsTokenValid(tokenString string) bool {
	_, err := ParseJwt(tokenString, nil, nil)
	return err == nil
}

func ParseJwt(tokenString string, salt interface{}, target interface{}) (interface{}, error) {

	claims, err := parseJwt(tokenString)
	if err != nil {
		return nil, err
	}

	if salt == nil {
		return claims, nil
	}

	info := (*claims)["info"]
	decryptString, err := tripledes.DecryptToString(info.(string), salt.(string))

	if err != nil {
		return nil, err
	}

	if (strings.HasPrefix(decryptString, "{") && strings.HasSuffix(decryptString, "}")) ||
		(strings.HasPrefix(decryptString, "[") && strings.HasSuffix(decryptString, "]")) {
		jsonErr := json.Unmarshal([]byte(decryptString), &target)
		if jsonErr != nil {
			return nil, err
		}
		(*claims)["info"] = target
	} else {
		(*claims)["info"] = decryptString
	}

	return claims, nil
}

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

func parseJwt(tokenString string) (*jwt.MapClaims, error) {

	globalConfig = config.GetGlobalConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(globalConfig.Application.Jwt.Key), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return &claims, nil
	} else {
		return nil, fmt.Errorf("Token invalid")
	}
}
