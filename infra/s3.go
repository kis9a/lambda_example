package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kis9a/lambda-sls/config"
)

var s3Uploader *s3manager.Uploader

func NewS3Uploader() {
	cfg := config.GetConfig()
	if cfg.ENV == "dev" {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(cfg.AWS_REGION),
			Credentials: credentials.NewStaticCredentials(
				cfg.AWS_ACCESS_KEY_ID, cfg.AWS_SECRET_ACCESS_KEY, "",
			),
		})
		if err != nil {
			return
		}
		s3Uploader = s3manager.NewUploader(sess)
	} else {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(cfg.AWS_REGION),
		}))
		s3Uploader = s3manager.NewUploader(sess)
	}
}

func GetS3Uploader() *s3manager.Uploader {
	return s3Uploader
}
