package backendGoPackage

import (
	"bytes"
	"mime/multipart"
	"net/url"

	"github.com/arutek/backend-go-package/pkg/db"
	"github.com/arutek/backend-go-package/pkg/helper"
	"github.com/arutek/backend-go-package/pkg/network"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)

func GormInit(host string, user string, pass string, name string, port string, sslMode string, tz string) *gorm.DB {
	return db.GormInit(host, user, pass, name, port, sslMode, tz)
}
func GormPaginate(page int, size int) func(paginate *gorm.DB) *gorm.DB {
	return db.PaginateDb(page, size)
}
func Error(err error, errCode string) map[string]interface{} {
	return helper.Error(err, errCode)
}
func SignAccessJwt(payload map[string]interface{}) (tokenString string, httpCode int, err error) {
	return helper.SignAccessJwt(payload)
}
func ValidateJwt(tokenString string) (token map[string]interface{}, httpCode int, err error) {
	return helper.ValidateJwt(tokenString)
}
func LoggerErr(msg string) {
	helper.LoggerErr(msg)
}
func LoggerWarn(msg string) {
	helper.LoggerWarn(msg)
}
func Logger(msg string) {
	helper.Logger(msg)
}
func AuthMiddleware(authHeader string) (token map[string]interface{}, errCode int, err error) {
	return helper.AuthMiddleware(authHeader)
}
func ParserCsv(file *multipart.FileHeader) (records [][]string, errRes map[string]interface{}) {
	return helper.ParserCsv(file)
}
func Response(msgVal string, dataVal interface{}, totalData int64) map[string]interface{} {
	return helper.Response(msgVal, dataVal, totalData)
}
func UploadS3Public(S3Client *s3.S3, bufferFile bytes.Buffer, uploadPath string, bucketName string, contentType string) (filepath string, err error) {
	return helper.UploadS3Public(S3Client, bufferFile, uploadPath, bucketName, contentType)
}
func UploadBuffer(uploadedFile *multipart.FileHeader) (file bytes.Buffer, fileName string, errRes map[string]interface{}) {
	return helper.UploadBuffer(uploadedFile)
}
func ApiGetData(urlString string, query map[string]string) ([]byte, error) {
	return network.ApiGetData(urlString, query)
}
func ApiPostData(urlString string, payload interface{}, query map[string]string) ([]byte, error) {
	return network.ApiPostData(urlString, payload, query)
}
func ApiPostForm(urlString string, formData url.Values, query map[string]string) ([]byte, error) {
	return network.ApiPostForm(urlString, formData, query)
}
