package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Add new reservation
func postReservation(w http.ResponseWriter, r *http.Request) {
	var resa Reservation
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&resa)

	log.Println(resa)
	if !isReservationValid(resa) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if existsReservationInRange(resa.Apartment_Id, resa.From, resa.To) {
		log.Printf("The reservation period matches with an existing one.")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = addReservation(resa)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// List all reservation
func getReservations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	reservations := findAllReservations()
	if len(*reservations) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reservations)
}

// List all reservation
func deleteReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	id := mux.Vars(r)["id"]

	if !deleteReservationById(id) {
		log.Printf("No reservation with id %s to delete", id)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	log.Printf("Reservation with id %s deleted", id)
	w.WriteHeader(http.StatusOK)
}
