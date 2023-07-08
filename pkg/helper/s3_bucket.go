package helper

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadS3Public(s3Client *s3.S3, bufferFile bytes.Buffer, uploadPath string, bucketName string, contentType string) (filepath string, err error) {
	filepath = ""
	object := s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(uploadPath),
		Body:        bytes.NewReader(bufferFile.Bytes()),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
	}
	_, err = s3Client.PutObject(&object)
	if err != nil {
		LoggerWarn(err.Error())
		return
	}
	filepath = "/" + bucketName + "/" + uploadPath
	return
}
