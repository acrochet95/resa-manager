package main

import (
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

// Main function initializing server with its HTTP endpoints
func main() {
	router := initializeRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
