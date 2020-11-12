package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Add new apartment
func postApartment(w http.ResponseWriter, r *http.Request) {
	var newApartment Apartment
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newApartment)

	if newApartment.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := addApartment(newApartment.Name)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// List all apartments
func getApartments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	apartments := getAllApartements()
	if len(*apartments) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apartments)
}

// Get one apartment by ID
func getApartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	id := mux.Vars(r)["id"]

	apartment := getApartementById(id)
	if apartment == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apartment)
}
