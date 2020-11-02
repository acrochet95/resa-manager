package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Add new apartment
func addApartment(w http.ResponseWriter, r *http.Request) {
	var newApartment apartment
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newApartment)

	if newApartment.ID == "" || newApartment.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	apartments = append(apartments, newApartment)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newApartment)
}

// List all apartments
func getApartments(w http.ResponseWriter, r *http.Request) {
	if len(apartments) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(apartments)
}

// Get one apartment by ID
func getApartment(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, apartment := range apartments {
		if apartment.ID == id {
			json.NewEncoder(w).Encode(apartment)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
