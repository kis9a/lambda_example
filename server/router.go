package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-todo/server/handlers"
)

func initRouter(r *gin.Engine) {
	r.GET("/health", handlers.HealthCheck())
	r.GET("/todos", handlers.GetTodos())
}
