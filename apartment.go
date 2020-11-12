package main

import (
	"log"
	"time"
)

// Apartment structure
type Apartment struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

type Apartments []Apartment

func getApartementById(id string) *Apartment {
	var apartment Apartment
	db := getDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM apartment WHERE a_id=$1", id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		err = rows.Scan(&apartment.ID, &apartment.Name, &apartment.Created_at)
		if err != nil {
			log.Fatal(err)
		}
		return &apartment
	}
	return nil
}

func getAllApartements() *Apartments {
	var apartments Apartments
	db := getDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM apartment")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var apartment Apartment
		rows.Scan(&apartment.ID, &apartment.Name, &apartment.Created_at)
		apartments = append(apartments, apartment)
	}

	return &apartments
}

func addApartment(name string) error {
	db := getDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO apartment (a_name, a_created_at) VALUES($1, $2)", name, time.Now())
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
