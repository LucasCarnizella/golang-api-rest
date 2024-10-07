package controllers

import (
	"encoding/json"
	"errors"
	"github.com/LucasCarnizella/golang-api-rest/database"
	"github.com/LucasCarnizella/golang-api-rest/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func ReadAllPersonas(w http.ResponseWriter, r *http.Request) {
	var personas []models.Persona
	var err error

	err = database.DB.Find(&personas).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(personas)
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
		return
	}

	if persona.Name == "" || persona.Story == "" {
		http.Error(w, "One of the parameters is an empty string", http.StatusBadRequest)
		return
	}

	err = database.DB.Create(&persona).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(persona)
	if err != nil {
		log.Panic(err)
		return
	}
}

func ReadPersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona
	var err error

	vars := mux.Vars(r)
	id := vars["id"]

	err = database.DB.First(&persona, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(persona)
	if err != nil {
		log.Panic(err)
		return
	}
}

func UpdatePersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona
	var err error

	vars := mux.Vars(r)
	id := vars["id"]

	err = database.DB.First(&persona, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		log.Panic(err)
		return
	}

	if persona.Name == "" || persona.Story == "" {
		http.Error(w, "One of the parameters is an empty string", http.StatusBadRequest)
		return
	}

	err = database.DB.Save(&persona).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona
	var err error

	vars := mux.Vars(r)
	id := vars["id"]

	err = database.DB.Delete(&persona, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Record not found", http.StatusNotFound)
		} else {
			http.Error(w, "", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
