package models

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kis9a/lambda-sls/infra"
)

type ImageUploader struct{}

func NewImageUploader() *ImageUploader {
	return &ImageUploader{}
}

const uploadImageBucket = "lambda-sls-images"

func (h *ImageUploader) Upload(file *multipart.FileHeader, fs multipart.File) (*s3manager.UploadOutput, error) {
	size := file.Size
	buffer := make([]byte, size)
	fs.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	uploader := infra.GetS3Uploader()
	uploadInput := &s3manager.UploadInput{
		Bucket:      aws.String(uploadImageBucket),
		Key:         aws.String(file.Filename),
		Body:        fileBytes,
		ACL:         aws.String("public-read"),
		ContentType: aws.String(fileType),
	}
	return uploader.Upload(uploadInput)
}
