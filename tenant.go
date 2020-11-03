package main

// Tenant structure
type tenant struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type Tenant []tenant

// tenant stored
var tenants Tenant
