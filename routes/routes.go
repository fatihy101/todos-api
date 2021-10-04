package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fatihy.space/todos-api/db"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes(db db.TodoDB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-User-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)
	r.Use(DBMiddleware(db))
	r.Use(JSONResponseMiddleware)
	r.Get("/", index)
	r.Route("/todos", todoRoutes)

	return r
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

// func AllowOriginFunc(r *http.Request, origin string) bool {
// 	if r.Method == "OPTIONS" {
// 		return true
// 	}
// 	return false
// }
