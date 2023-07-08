package helper

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func ParserCsv(file *multipart.FileHeader, fieldName string) (errRes map[string]interface{}, records [][]string) {
	src, err := file.Open()
	if err != nil {
		errRes = Error(err, "INVALID_CSV", http.StatusBadRequest)
		return
	}
	defer src.Close()
	if file.Header.Get("Content-Type") != "text/csv" {
		errRes = Error(err, "INVALID_HEADER_TYPE", http.StatusBadRequest)
		return
	}
	tempFile := fmt.Sprint("kodePoList_", time.Now().UnixMilli(), ".csv")
	dst, err := os.CreateTemp("", tempFile)
	if err != nil {
		errRes = Error(err, "CSV_FAIL", 0)
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
	csvFile, err := os.Open(dst.Name())
	if err != nil {
		errRes = Error(err, "CSV_FAIL", 0)
		return
	}
	defer csvFile.Close()
	records, err = csv.NewReader(csvFile).ReadAll()
	return
}
