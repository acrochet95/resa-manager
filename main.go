package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reservation manager server")
}

// Main function initializing server with its HTTP endpoints
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/apartments", getApartments).Methods("GET")
	router.HandleFunc("/apartments/{id}", getApartment).Methods("GET")
	router.HandleFunc("/apartments", postApartment).Methods("POST")

	router.HandleFunc("/tenants", getTenants).Methods("GET")
	router.HandleFunc("/tenants/{id}", getTenant).Methods("GET")
	router.HandleFunc("/tenants", postTenant).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
