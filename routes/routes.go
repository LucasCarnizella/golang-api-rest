package routes

import (
	"github.com/LucasCarnizella/golang-api-rest/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/ping", controllers.Ping)
	r.HandleFunc("/api/v1/personas", controllers.GetAllPersonas).Methods("Get")
	r.HandleFunc("/api/v1/personas/{id}", controllers.GetPersona).Methods("Get")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
