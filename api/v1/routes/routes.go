package routes

import (
	"github.com/gorilla/mux"
	"go-api/api/v1/handlers"
	"go-api/datastore"
	"net/http"
)

func InitialiseRoutes() *mux.Router {
	r := mux.NewRouter()
	app := handlers.LocationApp{
		DS: datastore.NewInMemeryStore(),
	}
	r.HandleFunc("/location/{user_id}", app.GetLocation).Methods("GET")
	r.HandleFunc("/location/{user_id}/now", app.NowLocation).Methods("POST")
	r.HandleFunc("/location/{user_id}", app.Delete).Methods("DELETE")
	r.Use(ContentTypeMiddleware)
	return r
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
