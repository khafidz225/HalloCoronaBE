package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoute(r)
	ArticleRoutes(r)
	ConsultationRoutes(r)
}
