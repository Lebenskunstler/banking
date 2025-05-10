package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city"  xml:"city"`
	Zipcode string `json:"zip_code"  xml:"zip-code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllcustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Alex", City: "New Jdida", Zipcode: "10000"},
		{Name: "Hassan", City: "New Sale", Zipcode: "2000"},
		{Name: "Madmax", City: "New york", Zipcode: "10000"},
	}

	if r.Header.Get("Content-Type") == "xml" {
		w.Header().Add("Content-Type", "xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "json")
		json.NewEncoder(w).Encode(customers)
	}

}
