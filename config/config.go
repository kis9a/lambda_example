package config

import (
	"os"
	"strconv"
)

type Config struct {
	ENV                   string
	AWS_S3_BUCKET         string
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
	AWS_REGION            string
	SERVER_PORT           string
	HTTP_AUTH_KEY         string
	HTTP_AUTH_SECRET      string
	DB_ENDPOINT           string
	DB_DISABLE_SSL        bool
}

var config Config

func NewConfig() {
	config.ENV = os.Getenv("ENV")
	config.SERVER_PORT = os.Getenv("SERVER_PORT")
	config.AWS_REGION = os.Getenv("AWS_REGION")
	config.AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
	config.AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	config.AWS_S3_BUCKET = os.Getenv("AWS_S3_BUCKET")
	config.HTTP_AUTH_KEY = os.Getenv("HTTP_AUTH_KEY")
	config.HTTP_AUTH_SECRET = os.Getenv("HTTP_AUTH_SECRET")
	config.DB_ENDPOINT = os.Getenv("DB_ENDPOINT")
	is, err := strconv.ParseBool(os.Getenv("DB_DISABLE_SSL"))
	if err != nil {
		is = false
	}
	config.DB_DISABLE_SSL = is
}

func GetConfig() *Config {
	return &config
}
