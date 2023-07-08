package helper

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadBuffer(ctx *gin.Context, fieldName string) (file bytes.Buffer, fileName string, err error, statusCode int) {
	statusCode = 0
	fileBuf, err := ctx.FormFile(fieldName)
	if err != nil {
		LoggerErr("Failed read form: " + err.Error())
		err = errors.New("INVALID_FILE")
		statusCode = http.StatusBadRequest
		return
	}
	fileName = fileBuf.Filename
	src, err := fileBuf.Open()
	if err != nil {
		LoggerErr("Failed buffer file: " + err.Error())
		err = errors.New("INVALID_FILE")
		statusCode = http.StatusBadRequest
		return
	}
	defer src.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(src)
	if err != nil {
		err = errors.New("FILE_FAIL")
		return
	}
	err = nil
	file = buffer
	return
}
