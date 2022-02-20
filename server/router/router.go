package router

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-sls/logger"
	"github.com/kis9a/lambda-sls/server/handlers"
	"github.com/kis9a/lambda-sls/server/middlewares"
)

func InitRouter(r *gin.Engine) {
	// get logger
	lg := logger.Getlogger()
	// init root middlewares
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(lg, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(lg, true))

	// Auth middleware
	r.Use(middlewares.AuthMiddleware())

	r.GET("/health", handlers.HealthCheck())
	r.GET("/todos", handlers.GetTodos())
}