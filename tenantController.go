package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Add new tenant
func postTenant(w http.ResponseWriter, r *http.Request) {
	var newTenant Tenant
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newTenant)

	if newTenant.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := addTenant(newTenant.Name)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// List all tenants
func getTenants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	tenants := getAllTenants()
	if len(*tenants) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenants)
}

// Get one tenant by ID
func getTenant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	id := mux.Vars(r)["id"]

	tenant := getTenantById(id)
	if tenant == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tenant)
}
