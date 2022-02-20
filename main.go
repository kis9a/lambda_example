package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kis9a/lambda-sls/config"
	"github.com/kis9a/lambda-sls/db"
	"github.com/kis9a/lambda-sls/logger"
	"github.com/kis9a/lambda-sls/server"
	"go.uber.org/zap"
)

func main() {
	// new config
	config.NewConfig()
	cfg := config.GetConfig()

	// init logger
	// new logger
	logger.Newlogger()

	// new db
	db.NewDB()

	s := server.NewServer()
	zap.L().Info("Listen Server Http ...",
		zap.String("port", s.Port),
		zap.String("mode", s.Mode),
	)
	if cfg.ENV == "dev" {
		err := s.ListenAndServeHttp()
		if err != nil {
			zap.L().Error("err", zap.Error(err))
		}
	} else {
		lambda.Start(s.ListenAndServeGinProxy)
	}
}
