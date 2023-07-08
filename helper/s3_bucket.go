package helper

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadS3Public(S3Client *s3.S3, bufferFile bytes.Buffer, uploadPath string, bucketName string, contentType string) (filepath string, err error) {
	filepath = ""
	object := s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(uploadPath),
		Body:        bytes.NewReader(bufferFile.Bytes()),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
	}
	_, err = S3Client.PutObject(&object)
	if err != nil {
		LoggerWarn(err.Error())
		return
	}
	filepath = os.Getenv("S3BUCKET_HOST") + "/" + bucketName + "/" + uploadPath
	return
}
