package models

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/kis9a/lambda-sls/infra"
)

const todoTable = "todo"

type Todo struct {
	TodoItems        []TodoItem `json:"todoItems"`
	LastEvaluatedKey TodoItem   `json:"lastEvaluatedKey"`
}

type TodoItem struct {
	Id   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func NewTodoItem() *Todo {
	return &Todo{}
}

func (t *Todo) CreateTodoItem(todo TodoItem) (TodoItem, error) {
	var err error

	ddb := infra.GetDB()
	id := uuid.New().String()

	todo.Id = id
	todo.Name = "apple"

	mmap, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return todo, err
	}

	params := &dynamodb.PutItemInput{
		Item:      mmap,
		TableName: aws.String(todoTable),
	}
	_, err = ddb.PutItem(params)
	return todo, err
}

func (t *Todo) ReadTodoItems(exclusiveStartKey TodoItem) (Todo, error) {
	var err error
	var todo Todo
	var todoItems []TodoItem

	ddb := infra.GetDB()

	startKeyMap, err := dynamodbattribute.MarshalMap(exclusiveStartKey)
	if err != nil {
		return todo, err
	}

	var params *dynamodb.ScanInput
	if exclusiveStartKey.Id != "" {
		params = &dynamodb.ScanInput{
			TableName:         aws.String(todoTable),
			Limit:             aws.Int64(20),
			ExclusiveStartKey: startKeyMap,
		}
	} else {
		params = &dynamodb.ScanInput{
			TableName: aws.String(todoTable),
			Limit:     aws.Int64(20),
		}
	}
	scan, err := ddb.Scan(params)

	for _, i := range scan.Items {
		t := TodoItem{}
		err = dynamodbattribute.UnmarshalMap(i, &t)
		if err != nil {
			return todo, err
		}
		todoItems = append(todoItems, t)
	}

	var LastEvaluatedKey TodoItem
	dynamodbattribute.UnmarshalMap(scan.LastEvaluatedKey, &LastEvaluatedKey)

	todo.TodoItems = todoItems
	todo.LastEvaluatedKey = LastEvaluatedKey
	return todo, err
}

func (t *Todo) DeleteTodoItem(todo TodoItem) (TodoItem, error) {
	ddb := infra.GetDB()
	keyMap, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return todo, err
	}
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(todoTable),
		Key:       keyMap,
	}
	_, err = ddb.DeleteItem(params)
	return todo, err
}

func (t *Todo) UpdateTodoItem(todo TodoItem) (TodoItem, error) {
	ddb := infra.GetDB()
	keyMap, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return todo, err
	}
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(todoTable),
		Key:       keyMap,
	}
	_, err = ddb.UpdateItem(params)
	return todo, err
}
