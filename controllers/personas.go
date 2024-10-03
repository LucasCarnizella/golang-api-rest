package controllers

import (
	"encoding/json"
	"github.com/LucasCarnizella/golang-api-rest/database"
	"github.com/LucasCarnizella/golang-api-rest/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ReadAllPersonas(w http.ResponseWriter, r *http.Request) {
	var personas []models.Persona

	database.DB.Find(&personas)

	err := json.NewEncoder(w).Encode(personas)
	if err != nil {
		log.Panic(err)
	}
}

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona
	var err error

	err = json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		log.Panic(err)
	}

	if persona.Name == "" || persona.Story == "" {
		log.Panic("One of the variables is an empty string.")
	}

	database.DB.Create(&persona)

	err = json.NewEncoder(w).Encode(persona)
	if err != nil {
		log.Panic(err)
	}
}

func ReadPersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona

	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&persona, id)
	err := json.NewEncoder(w).Encode(persona)
	if err != nil {
		log.Panic(err)
	}
}

func UpdatePersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona
	var err error

	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&persona, id)

	err = json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		log.Panic(err)
	}

	if persona.Name == "" || persona.Story == "" {
		log.Panic("One of the variables is an empty string.")
	}

	database.DB.Save(&persona)

	err = json.NewEncoder(w).Encode(&persona)
	if err != nil {
		log.Panic(err)
	}
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona

	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&persona, id)

	err := json.NewEncoder(w).Encode(persona)
	if err != nil {
		log.Panic(err)
	}

	database.DB.Delete(&persona, id)
}
