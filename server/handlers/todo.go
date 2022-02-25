package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-sls/models"
	"go.uber.org/zap"
)

func ReadTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.NewTodoItem()
		var todoItem models.TodoItem
		c.ShouldBindJSON(&todoItem)
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil && limit == 0 {
			limit = 20
		}
		todo, err := model.ReadTodoItems(todoItem, int64(limit))
		if err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todo})
	}
}

func CreateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.NewTodoItem()
		var todoItem models.TodoItem
		if err := c.ShouldBindJSON(&todoItem); err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo, err := model.CreateTodoItem(todoItem)
		if err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todo})
	}
}

func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.NewTodoItem()
		var todoItem models.TodoItem
		if err := c.ShouldBindJSON(&todoItem); err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo, err := model.DeleteTodoItem(todoItem)
		if err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todo})
	}
}

func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.NewTodoItem()
		var todoItem models.TodoItem
		if err := c.ShouldBindJSON(&todoItem); err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo, err := model.UpdateTodoItem(todoItem)
		if err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todo})
	}
}
