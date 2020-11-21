package main

import (
	"log"
	"time"
)

// Apartment structure
type Reservation struct {
	ID           int       `json:"id"`
	Apartment_Id int       `json:"apartment_id,omitempty"`
	Tenant_Id    int       `json:"tenant_id,omitempty"`
	Price        int       `json:"price"`
	Description  string    `json:"description"`
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
	Paid         bool      `json:"isPaid"`
}

type Reservations []Reservation

func findAllReservations() *Reservations {
	var reservations Reservations
	db := getDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM reservation")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var resa Reservation
		err = rows.Scan(&resa.ID, &resa.Apartment_Id, &resa.Tenant_Id, &resa.Price, &resa.Description, &resa.From, &resa.To, &resa.Paid)
		if err != nil {
			log.Fatal(err)
		}
		reservations = append(reservations, resa)
	}

	return &reservations
}

func existsReservationInRange(apartment_id int, startDate time.Time, endDate time.Time) bool {
	db := getDatabase()
	defer db.Close()

	var count int
	rows := db.QueryRow("SELECT COUNT(*) as count FROM reservation WHERE a_id=$1 AND (($2 >= r_from AND $2 <= r_to) OR ($3 >= r_from AND $3 <= r_to) OR ($2 <= r_from AND $3 >= r_to))",
		apartment_id, startDate, endDate)
	rows.Scan(&count)

	return count > 0
}

func addReservation(resa Reservation) error {
	db := getDatabase()
	_, err := db.Exec("INSERT INTO reservation (a_id, t_id, r_price, r_description, r_from, r_to, r_paid) VALUES ($1,$2,$3,$4,$5,$6,$7)",
		resa.Apartment_Id, resa.Tenant_Id, resa.Price, resa.Description, resa.From, resa.To, resa.Paid)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func isReservationValid(resa Reservation) bool {
	return resa.Apartment_Id != 0 && resa.Tenant_Id != 0 && resa.Price != 0 && !resa.From.IsZero() && !resa.To.IsZero() && (resa.From.Before(resa.To) || resa.From.Equal(resa.To))
}
