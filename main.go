package main

import (
	"fmt"
	"github.com/LucasCarnizella/golang-api-rest/database"
	"github.com/LucasCarnizella/golang-api-rest/routes"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")

	fmt.Println("Starting Server on Port " + port)

	database.DbConnect()
	routes.HandleRequest(port)
}
