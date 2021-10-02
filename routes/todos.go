package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"fatihy.space/todos-api/db"
	"github.com/go-chi/chi/v5"
)

var (
	successMap         = map[string]interface{}{"success": true}
	errorMap           = map[string]interface{}{"success": false}
	ErrUserIdMismatch  = errors.New("validation failed: Header userId and item's userId must be same")
	ErrNoTodosFound    = errors.New("todos not found")
	ErrUserIdRequired  = errors.New("validation failed: Header userId is required")
	ErrrTodoIdRequired = errors.New("validation failed: id is required")
	ErrNoTodoFound     = errors.New("todo item couldn't found with the given id")
)

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-User-Id")
	mdb := getDB(r)
	todos, err := mdb.GetTodos(userId)

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", ErrUserIdRequired.Error())
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", err.Error())
		return
	}

	if len(todos) == 0 {
		w.WriteHeader(http.StatusNotFound)
		createResponse(w, errorMap, "error", ErrNoTodosFound.Error())
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-User-Id")
	mdb := getDB(r)
	todo := db.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", err.Error())
		return
	}

	if todo.UserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		createResponse(w, errorMap, "error", ErrUserIdMismatch.Error())
		return
	}

	err = mdb.AddTodo(userId, todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", err.Error())
		return
	}

	createResponse(w, successMap, nil, nil)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-User-Id")
	mdb := getDB(r)
	itemId := chi.URLParam(r, "id")

	item, err := mdb.GetTodoById(itemId)

	if err != nil {
		fmt.Printf("Error on delete handler: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", ErrNoTodoFound.Error())
		return
	}

	if item.UserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		createResponse(w, errorMap, "error", ErrUserIdMismatch.Error())
		return
	}

	err = mdb.DeleteTodo(itemId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", err.Error())
		return
	}

	createResponse(w, successMap, nil, nil)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("X-User-Id")
	mdb := getDB(r)
	todo := db.Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", err.Error())
		return
	}

	if todo.UserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		createResponse(w, errorMap, "error", ErrUserIdMismatch.Error())
		return
	}

	err = mdb.UpdateTodo(userId, todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		createResponse(w, errorMap, "error", err.Error())
		return
	}

	createResponse(w, successMap, nil, nil)
}
