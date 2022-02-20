package handlers

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-sls/db"
	"github.com/kis9a/lambda-sls/models"
	"go.uber.org/zap"
)

type Movie struct {
	Year  int    `json:"year" dynamodbav:"year"`
	Title string `json:"title" dynamodbav:"title"`
	St    string `json:"st" dynamodbav:"st"`
	// Info  MovieInfo `json:"info" dynamodbav:"info"`
}

func GetTodos() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todos, err := models.Todos().AllG(c)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		// 	return
		// }
		ddb := db.GetDB()

		movie := Movie{
			Year:  2016,
			Title: "The Big New Moviee",
			St:    "apple",
		}

		mmap, err := dynamodbattribute.MarshalMap(movie)
		if err != nil {
			panic("Cannot marshal movie into AttributeValue map")
		}

		// ddbTable, err := ddb.CreateTable(&dynamodb.CreateTableInput{
		// 	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		// 		{
		// 			AttributeName: aws.String("id"), // プライマリキー名
		// 			AttributeType: aws.String("S"),  // データ型(String:S, Number:N, Binary:B の三種)
		// 		},
		// 	},
		// 	KeySchema: []*dynamodb.KeySchemaElement{
		// 		{
		// 			AttributeName: aws.String("id"),   // インデックス名
		// 			KeyType:       aws.String("HASH"), // インデックスの型(HASH または RANGE)
		// 		},
		// 	},
		// 	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
		// 		ReadCapacityUnits:  aws.Int64(1), // 読み込みキャパシティーユニット（デフォルト：１）
		// 		WriteCapacityUnits: aws.Int64(1), // 書き込みキャパシティーユニット（デフォルト：１）
		// 	},
		// 	TableName: aws.String("Demo"), // テーブル名
		// })
		// zap.S().Info(ddbTable)

		// create the api params
		params := &dynamodb.PutItemInput{
			TableName: aws.String("Movies"),
			Item:      mmap,
		}

		res, err := ddb.PutItem(params)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
		}
		zap.S().Info(res)

		c.JSON(http.StatusOK, gin.H{"todos": []models.Todo{}})
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
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": models.Todo{}})
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
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": models.Todo{}})
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
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": models.Todo{}})
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
		c.JSON(http.StatusOK, gin.H{"success": true, "todo": models.Todo{}})
	}
}
