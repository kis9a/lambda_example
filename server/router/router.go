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
	r.POST("/health", handlers.HealthCheck())

	// todo router
	r.POST("/todos", handlers.ReadTodos())
	r.POST("/todos/create", handlers.CreateTodo())
	r.POST("/todos/update", handlers.UpdateTodo())
	r.POST("/todos/delete", handlers.DeleteTodo())
}
