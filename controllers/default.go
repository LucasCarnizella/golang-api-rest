package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/LucasCarnizella/golang-api-rest/database"
	"github.com/LucasCarnizella/golang-api-rest/models"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "")
	if err != nil {
		log.Panic(err)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	var utils models.Utils

	database.DB.Find(&utils, "key = ?", "ping")
	//_, err := fmt.Fprintf(w, utils.Value)
	//if err != nil {
	//	log.Panic(err)
	//}
	err := json.NewEncoder(w).Encode(utils.Value)
	if err != nil {
		log.Panic(err)
	}
}
