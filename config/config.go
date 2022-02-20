package config

import "os"

type Config struct {
	ENV                   string
	AWS_S3_BUCKET         string
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
	AWS_REGION            string
	SERVER_PORT           string
}

func NewConfig() *Config {
	return &Config{
		ENV:                   os.Getenv("ENV"),
		SERVER_PORT:           os.Getenv("SERVER_PORT"),
		AWS_REGION:            os.Getenv("AWS_REGION"),
		AWS_ACCESS_KEY_ID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWS_SECRET_ACCESS_KEY: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AWS_S3_BUCKET:         os.Getenv("AWS_S3_BUCKET"),
	}
}
