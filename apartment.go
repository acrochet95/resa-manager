package main

// Apartment structure
type apartment struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type Apartments []apartment

// apartments stored
var apartments Apartments
