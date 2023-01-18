package routes

import (
	handlers "server/handler"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func UserRoute(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUser).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	// r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
}
