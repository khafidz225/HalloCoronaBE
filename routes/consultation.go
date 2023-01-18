package routes

import (
	handlers "server/handler"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gorilla/mux"
)

func ConsultationRoutes(r *mux.Router) {
	consultationRepository := repositories.RepositoryConsultation(mysql.DB)
	h := handlers.HandlerConsultation(consultationRepository)

	r.HandleFunc("/consultations", h.FindConsultation).Methods("GET")
	r.HandleFunc("/consultations/{id}", h.GetConsultation).Methods("GET")
	r.HandleFunc("/consultations", middleware.Auth(h.CreateConsultation)).Methods("POST")
	r.HandleFunc("/consultations/{id}", h.UpdateConsultation).Methods("PATCH")
}
