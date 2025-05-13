package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

// FindAll retrieves all customers from the database
// and returns them as a slice of Customer structs.
// It returns an error if there is an issue with the database query.
func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.ZipCode, &customer.DateofBirth, &customer.Status)
		if err != nil {
			log.Println("Error while scanning customer row " + err.Error())
			return nil, err
		}

		customers = append(customers, customer)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:toor@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
