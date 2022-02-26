package models

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/kis9a/lambda-sls/config"
	"github.com/kis9a/lambda-sls/infra"
	"github.com/pkg/errors"
)

type Todo struct {
	TodoTable string `json:"todoTable"`
}

type TodoItemRequest struct {
	Id        string                     `json:"id" dynamodbav:"id"`
	Name      string                     `json:"name" dynamodbav:"name"`
	CreatedAt dynamodbattribute.UnixTime `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt dynamodbattribute.UnixTime `json:"updated_at" dynamodbav:"updated_at"`
}

type TodoItemResponse struct {
	Id        string    `json:"id" dynamodbav:"id"`
	Name      string    `json:"name" dynamodbav:"name"`
	CreatedAt time.Time `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt time.Time `json:"updated_at" dynamodbav:"updated_at"`
}

type TodoReadReponse struct {
	TodoItems []TodoItemResponse `json:"todoItems"`
	Next      TodoItemRequest    `json:"next"`
}

type TodoReadRequest struct {
	Limit     int64           `json:"limit"`
	StartItem TodoItemRequest `json:"start_at"`
}

type TodoItemCreateRequest struct {
	Name string `json:"name"`
}

func NewTodoItem() *Todo {
	cfg := config.GetConfig()
	return &Todo{
		TodoTable: cfg.DB_TODO_TABLE,
	}
}

func (t *Todo) CreateTodoItem(req TodoItemCreateRequest) (TodoItemResponse, error) {
	var todo TodoItemRequest
	var todoRes TodoItemResponse
	var err error

	ddb := infra.GetDB()

	count := 0
	isAlreadExists := true
	for isAlreadExists {
		todo.Id = uuid.New().String()
		getParams := &dynamodb.GetItemInput{
			TableName: aws.String(t.TodoTable),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(todo.Id),
				},
			},
		}
		get, err := ddb.GetItem(getParams)
		if err != nil {
			return todoRes, err
		}

		if len(get.Item) == 0 {
			isAlreadExists = false
		} else {
			// already existed Id
			if count > 5 {
				return todoRes, errors.Errorf("can't uniq Todo Item Id")
			}
		}
		count++
	}
	todo.Name = req.Name
	todo.CreatedAt = dynamodbattribute.UnixTime(time.Now())
	todo.UpdatedAt = dynamodbattribute.UnixTime(time.Now())

	mmap, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return todoRes, err
	}

	params := &dynamodb.PutItemInput{
		Item:      mmap,
		TableName: aws.String(t.TodoTable),
	}

	_, err = ddb.PutItem(params)
	if err != nil {
		return todoRes, err
	}

	todoRes = TodoItemResponse{
		Id:        todo.Id,
		Name:      todo.Name,
		CreatedAt: time.Time(todo.CreatedAt),
		UpdatedAt: time.Time(todo.UpdatedAt),
	}
	return todoRes, err
}

func (t *Todo) ReadTodoItems(req TodoReadRequest) (TodoReadReponse, error) {
	var err error
	var todoRes TodoReadReponse

	ddb := infra.GetDB()

	startKeyMap, err := dynamodbattribute.MarshalMap(req.StartItem)
	if err != nil {
		return todoRes, err
	}

	var params *dynamodb.ScanInput
	if req.StartItem.Id != "" {
		params = &dynamodb.ScanInput{
			TableName:         aws.String(t.TodoTable),
			Limit:             aws.Int64(req.Limit),
			ExclusiveStartKey: startKeyMap,
		}
	} else {
		params = &dynamodb.ScanInput{
			TableName: aws.String(t.TodoTable),
			Limit:     aws.Int64(req.Limit),
		}
	}
	scan, err := ddb.Scan(params)

	var todoResItems []TodoItemResponse
	for _, i := range scan.Items {
		item := TodoItemRequest{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return todoRes, err
		}
		todoResItems = append(todoResItems, TodoItemResponse{
			Id:        item.Id,
			Name:      item.Name,
			CreatedAt: time.Time(item.CreatedAt),
			UpdatedAt: time.Time(item.UpdatedAt),
		})
	}

	var LastEvaluatedKey TodoItemRequest
	dynamodbattribute.UnmarshalMap(scan.LastEvaluatedKey, &LastEvaluatedKey)

	todoRes.TodoItems = todoResItems
	todoRes.Next = LastEvaluatedKey
	return todoRes, err
}

func (t *Todo) DeleteTodoItem(todo TodoItemRequest) (TodoItemResponse, error) {
	var todoRes TodoItemResponse

	ddb := infra.GetDB()

	keyMap, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return todoRes, err
	}

	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(t.TodoTable),
		Key:       keyMap,
	}

	_, err = ddb.DeleteItem(params)
	if err != nil {
		return todoRes, err
	}

	todoRes = TodoItemResponse{
		Id:        todo.Id,
		Name:      todo.Name,
		CreatedAt: time.Time(todo.CreatedAt),
		UpdatedAt: time.Time(todo.UpdatedAt),
	}
	return todoRes, err
}

func (t *Todo) UpdateTodoItem(todo TodoItemRequest) (TodoItemResponse, error) {
	var todoRes TodoItemResponse

	ddb := infra.GetDB()

	todo.UpdatedAt = dynamodbattribute.UnixTime(time.Now())

	keyMap, err := dynamodbattribute.MarshalMap(todo)
	if err != nil {
		return todoRes, err
	}

	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(t.TodoTable),
		Key:       keyMap,
	}
	_, err = ddb.UpdateItem(params)
	if err != nil {
		return todoRes, err
	}

	todoRes = TodoItemResponse{
		Id:        todo.Id,
		Name:      todo.Name,
		CreatedAt: time.Time(todo.CreatedAt),
		UpdatedAt: time.Time(todo.UpdatedAt),
	}
	return todoRes, err
}
