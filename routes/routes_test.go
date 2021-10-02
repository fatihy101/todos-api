package routes

import (
	"bytes"
	"encoding/json"
	"errors"

	"fatihy.space/todos-api/db"
)

var (
	ErrUnexpectedStatus  = errors.New("unexpected status code")
	ErrUnexpectedResBody = errors.New("unexpected response body")
)

type mockDB struct {
	todos []db.Todo
}

func (m *mockDB) GetTodos(userId string) ([]db.Todo, error) {
	return m.todos, nil
}

func (m *mockDB) DeleteTodo(todoId string) error {
	for i, todoItem := range m.todos {
		if todoItem.ID == todoId {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *mockDB) UpdateTodo(userID string, todo db.Todo) error {
	for i, todoItem := range m.todos {
		if todoItem.ID == todo.ID {
			m.todos[i] = todo
			return nil
		}
	}
	return nil
}

func (m *mockDB) AddTodo(userID string, todo db.Todo) error {
	m.todos = append(m.todos, todo)
	return nil
}

func (m *mockDB) GetTodoById(todoId string) (db.Todo, error) {
	for _, todoItem := range m.todos {
		if todoItem.ID == todoId {
			return todoItem, nil
		}
	}
	return db.Todo{}, nil
}

func toJsonStr(val interface{}) string {
	returnVal, _ := json.Marshal(val)
	return string(returnVal)
}

func toJson(val interface{}) *bytes.Reader {
	returnVal, _ := json.Marshal(val)
	return bytes.NewReader(returnVal)
}
