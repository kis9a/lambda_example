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
		var todoItem models.TodoItemRequest
		c.ShouldBindJSON(&todoItem)
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil && limit == 0 {
			limit = 20
		}
		todo, err := model.ReadTodoItems(models.TodoReadRequest{
			Limit:     int64(limit),
			StartItem: todoItem,
		})
		if err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.NewTodoItem()
		var createItem models.TodoItemCreateRequest
		if err := c.ShouldBindJSON(&createItem); err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo, err := model.CreateTodoItem(createItem)
		if err != nil {
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.NewTodoItem()
		var todoItem models.TodoItemRequest
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
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		model := models.NewTodoItem()
		var todoItem models.TodoItemRequest
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
		c.JSON(http.StatusOK, todo)
	}
}
