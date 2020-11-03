package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Add new tenant
func addTenant(w http.ResponseWriter, r *http.Request) {
	var newTenant tenant
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newTenant)

	if newTenant.ID == "" || newTenant.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tenants = append(tenants, newTenant)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTenant)
}

// List all tenants
func getTenants(w http.ResponseWriter, r *http.Request) {
	if len(tenants) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tenants)
}

// Get one tenant by ID
func getTenant(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, tenant := range tenants {
		if tenant.ID == id {
			json.NewEncoder(w).Encode(tenant)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
