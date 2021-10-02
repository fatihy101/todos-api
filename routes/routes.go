package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fatihy.space/todos-api/db"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(db db.TodoDB) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(AllowOrigin)
	router.Use(DBMiddleware(db))
	router.Use(JSONResponseMiddleware)
	router.Get("/", index)
	router.Route("/todos", todoRoutes)

	return router
}

func todoRoutes(r chi.Router) {
	r.Get("/", GetTodosHandler)
	r.Post("/", AddTodoHandler)
	r.Delete("/{id}", DeleteTodoHandler)
	r.Put("/", UpdateTodoHandler)
}

func index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "Server is active"})
}

func createResponse(w http.ResponseWriter, baseMap map[string]interface{}, key interface{}, value interface{}) {
	if value != nil {
		baseMap[fmt.Sprintf("%v", key)] = value
	}
	json.NewEncoder(w).Encode(baseMap)
}
