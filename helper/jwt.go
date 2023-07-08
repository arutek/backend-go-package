package helper

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignAccessJwt(payload map[string]interface{}) (tokenString string, httpCode int, err error) {
	httpCode = 0
	payload["expiredAt"] = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	tokenString, err = token.SignedString([]byte(os.Getenv("MIDDLEWARE_SECRET")))
	if err != nil {
		LoggerErr(err.Error())
		err = fmt.Errorf("JWT_FAIL")
	}
	return
}
func ValidateJwt(tokenString string) (token map[string]interface{}, httpCode int, err error) {
	httpCode = http.StatusUnauthorized
	tokenPointer, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("MIDDLEWARE_SECRET")), nil
	})
	if err != nil {
		err = fmt.Errorf("TOKEN_INVALID")
		return
	}
	token, ok := tokenPointer.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("TOKEN_INVALID")
	}
	return
}
