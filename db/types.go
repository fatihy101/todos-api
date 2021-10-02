package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoDB interface {
	GetTodos(userID string) ([]Todo, error)
	DeleteTodo(todoId string) error
	UpdateTodo(userID string, todo Todo) error
	AddTodo(userID string, todo Todo) error
	GetTodoById(todoId string) (Todo, error)
}

type Todo struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Text      string `json:"text" bson:"text"`
	Completed bool   `json:"completed" bson:"completed"`
	UserId    string `json:"userId" bson:"userId"`
}

type DBHandle struct {
	monDB *mongo.Database
}
