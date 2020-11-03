package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/apartments", addApartment).Methods("POST")

	router.HandleFunc("/tenants", getTenants).Methods("GET")
	router.HandleFunc("/tenants/{id}", getTenant).Methods("GET")
	router.HandleFunc("/tenants", addTenant).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
