package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kis9a/lambda-todo/config"
	"github.com/kis9a/lambda-todo/server"
	"go.uber.org/zap"
)

func main() {
	// new config
	cfg := config.NewConfig()

	// init logger
	// new logger
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	// serve
	s := server.NewServer(cfg, logger)
	if cfg.ENV == "dev" {
		err := s.ListenAndServeHttp()
		if err != nil {
			zap.L().Error(err.Error())
		}
	} else {
		lambda.Start(s.ListenAndServeGinProxy)
	}
}
