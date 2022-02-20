package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HandlerTodo struct{}

type Todo struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

func GetTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todos, err := models.Todos().AllG(c)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{"todos": []Todo{}})
	}
}

func GetTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// type Data struct {
		// 	Id int
		// }
		// var data Data
		// if err := c.ShouldBindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// todo, err := models.Todos(
		// 	qm.Where("id=?", data.Id),
		// ).OneG(c)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": Todo{}})
	}
}

func PostTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var todo models.Todo
		// if err := c.ShouldBindJSON(&todo); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// err := todo.InsertG(c, boil.Infer())
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		// 	log.Println(err)
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": Todo{}})
	}
}

func UpdateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var todo models.Todo
		// if err := c.ShouldBindJSON(&todo); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// _, err := todo.UpdateG(c, boil.Infer())
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		// }
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": Todo{}})
	}
}

func DeleteTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// type Data struct {
		// 	Id int64
		// }
		// var data Data
		// if err := c.ShouldBindJSON(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// var todo models.Todo
		// todo.ID = data.Id
		// _, err := todo.DeleteG(c)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": Todo{}})
	}
}
