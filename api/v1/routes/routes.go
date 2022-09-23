package routes

import (
	"github.com/gorilla/mux"
	"go-api/api/v1/handlers"
	"net/http"
)

func InitialiseRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/demo/{type}", handlers.Demofunc).Methods("GET")
	r.Use(AuthMiddleware)
	return r
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.SetBasicAuth("innoe", "p@ssword")

		next.ServeHTTP(w, r)
	})
}
