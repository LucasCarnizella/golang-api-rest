package routes

import (
	"github.com/LucasCarnizella/golang-api-rest/controllers"
	"github.com/LucasCarnizella/golang-api-rest/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest(port string) {
	r := mux.NewRouter()

	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/ping", controllers.Ping)
	r.HandleFunc("/api/v1/personas", controllers.ReadAllPersonas).Methods("Get")
	r.HandleFunc("/api/v1/personas", controllers.CreatePersona).Methods("Post")
	r.HandleFunc("/api/v1/personas/{id}", controllers.ReadPersona).Methods("Get")
	r.HandleFunc("/api/v1/personas/{id}", controllers.UpdatePersona).Methods("Put")
	r.HandleFunc("/api/v1/personas/{id}", controllers.DeletePersona).Methods("Delete")

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
