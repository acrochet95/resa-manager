package main

import "github.com/gorilla/mux"

func initializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// endpoints for apartment entity
	router.HandleFunc("/apartments", getApartments).Methods("GET")
	router.HandleFunc("/apartments/{id}", getApartment).Methods("GET")
	router.HandleFunc("/apartments", postApartment).Methods("POST")

	// endpoints for tenant entity
	router.HandleFunc("/tenants", getTenants).Methods("GET")
	router.HandleFunc("/tenants/{id}", getTenant).Methods("GET")
	router.HandleFunc("/tenants", postTenant).Methods("POST")

	// endpoints for reservation entity
	router.HandleFunc("/reservations", getReservations).Methods("GET")
	router.HandleFunc("/reservations", postReservation).Methods("POST")
	router.HandleFunc("/reservations/{id}", deleteReservation).Methods("DELETE")

	return router
}
