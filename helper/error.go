package helper

import (
	"net/http"
)

func Error(err error, errCode string, statusCode int) map[string]interface{} {
	if statusCode <= 0 {
		statusCode = http.StatusInternalServerError
	}
	LoggerErr(err.Error())
	return Response(errCode, nil, -1)
}
