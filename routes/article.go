package routes

import (
	handlers "server/handler"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func ArticleRoutes(r *mux.Router) {
	articleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(articleRepository)

	r.HandleFunc("/article", h.FindArticle).Methods("GET")
	r.HandleFunc("/article/{id}", h.GetArticle).Methods("GET")
	r.HandleFunc("/article", middleware.Auth(middleware.UploadFile(h.AddArticle))).Methods("POST")
	r.HandleFunc("/article/{id}", middleware.Auth(middleware.UploadFile(h.UpdateArticle))).Methods("PATCH")
	r.HandleFunc("/article/{id}", h.DeleteArticle).Methods("DELETE")
}
