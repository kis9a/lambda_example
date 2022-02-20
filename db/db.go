package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kis9a/lambda-sls/config"
)

var db *dynamodb.DynamoDB

func NewDB() {
	cfg := config.GetConfig()
	if cfg.ENV == "dev" {
		db = dynamodb.New(session.New(&aws.Config{
			Region:   aws.String(cfg.AWS_REGION),
			Endpoint: aws.String(cfg.DB_ENDPOINT),
		}))
	} else {
		db = dynamodb.New(session.New(&aws.Config{
			Region:      aws.String(cfg.AWS_REGION),
			Credentials: credentials.NewEnvCredentials(),
			Endpoint:    aws.String(cfg.DB_ENDPOINT),
			DisableSSL:  aws.Bool(cfg.DB_DISABLE_SSL),
		}))
	}
}

func GetDB() *dynamodb.DynamoDB {
	return db
}
