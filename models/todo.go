package models

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/kis9a/lambda-sls/db"
	"go.uber.org/zap"
)

type Todo struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
}

func (h Todo) GetTodos(id string) (*Todo, error) {
	db := db.GetDB()
	params := &dynamodb.GetItemInput{}
	resp, err := db.GetItem(params)
	if err != nil {
		zap.L().Error("error db get item", zap.Error(err))
		return nil, err
	}
	var todo *Todo
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &todo); err != nil {
		zap.L().Error("error db get item", zap.Error(err))
		return nil, err
	}
	return todo, nil
}

func (h Todo) PostTodos(id string) (*Todo, error) {
	db := db.GetDB()
	params := &dynamodb.GetItemInput{}
	resp, err := db.GetItem(params)
	if err != nil {
		zap.L().Error("error db get item", zap.Error(err))
		return nil, err
	}
	var todo *Todo
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &todo); err != nil {
		zap.L().Error("error db get item", zap.Error(err))
		return nil, err
	}
	return todo, nil
}
