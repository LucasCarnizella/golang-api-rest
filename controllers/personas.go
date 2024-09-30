package controllers

import (
	"encoding/json"
	"github.com/LucasCarnizella/golang-api-rest/database"
	"github.com/LucasCarnizella/golang-api-rest/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetAllPersonas(w http.ResponseWriter, r *http.Request) {
	var personas []models.Persona

	database.DB.Find(&personas)
	err := json.NewEncoder(w).Encode(personas)
	if err != nil {
		log.Panic(err)
	}
}

func GetPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var persona models.Persona

	database.DB.First(&persona, id)
	err := json.NewEncoder(w).Encode(persona)
	if err != nil {
		log.Panic(err)
	}
}
