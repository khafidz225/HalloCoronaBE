package main

import (
	"fmt"
	"net/http"
	"os"
	"server/database"
	"server/pkg/mysql"
	"server/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	errNev := godotenv.Load()
	if errNev != nil {
		fmt.Println(errNev)
	}

	mysql.DatabaseInit()
	database.RunMigration()
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	AllowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	AllowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS", "HEAD"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = os.Getenv("PORT")

	fmt.Println("Server runing localhost:" + port)
	// http.ListenAndServe("localhost:5000", r)
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
