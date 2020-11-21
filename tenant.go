package main

import (
	"log"
	"time"
)

// Tenant structure
type Tenant struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

type Tenants []Tenant

func getTenantById(id string) *Tenant {
	var tenant Tenant
	db := getDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tenant WHERE t_id=$1", id)
	defer rows.Close()
	if err != nil {
		log.Panic(err)
	}
	if rows.Next() {
		err = rows.Scan(&tenant.ID, &tenant.Name, &tenant.Created_at)
		if err != nil {
			log.Panic(err)
		}
		return &tenant
	}
	return nil
}

func getAllTenants() *Tenants {
	var tenants Tenants
	db := getDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tenant")
	defer rows.Close()
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var tenant Tenant
		rows.Scan(&tenant.ID, &tenant.Name, &tenant.Created_at)
		tenants = append(tenants, tenant)
	}

	return &tenants
}

func addTenant(name string) error {
	db := getDatabase()
	defer db.Close()

	_, err := db.Exec("INSERT INTO tenant (t_name, t_created_at) VALUES($1, $2)", name, time.Now())
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
