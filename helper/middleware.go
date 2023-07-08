package helper

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(authHeader string) (token map[string]interface{}, err error, errCode int) {
	authHeaders := strings.Split(authHeader, " ")
	if len(authHeader) < 2 {
		err = errors.New("UNAUTHORIZED")
		errCode = http.StatusUnauthorized
		return
	}
	claims, err, statusCode := ValidateJwt(authHeaders[1])
	if err != nil {
		errCode = statusCode
		return
	}
	if int64(claims["expiredAt"].(float64)) < time.Now().Unix() {
		err = errors.New("TOKEN_EXPIRED")
		errCode = statusCode
		return
	}
	errCode = 0
	err = nil
	token = claims
	return
}
