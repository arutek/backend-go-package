package helper

import (
	"bytes"
	"mime/multipart"
)

func UploadBuffer(uploadedFile *multipart.FileHeader) (file bytes.Buffer, fileName string, errRes map[string]interface{}) {
	fileName = uploadedFile.Filename
	src, err := uploadedFile.Open()
	if err != nil {
		LoggerErr("Failed buffer file: " + err.Error())
		errRes = Error(err, "INVALID_FILE")
		return
	}
	defer src.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(src)
	if err != nil {
		errRes = Error(err, "FILE_FAIL")
		return
	}
	err = nil
	file = buffer
	return
}
