package routes

import (
	"context"

	"net/http"

	"fatihy.space/todos-api/db"
)

type key int

const (
	DBContext key = iota
)

func DBMiddleware(db db.TodoDB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), DBContext, db)))
		})
	}
}

func JSONResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func getDB(r *http.Request) db.TodoDB {
	return r.Context().Value(DBContext).(db.TodoDB)
}
