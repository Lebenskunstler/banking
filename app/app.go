package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/orgs/Lebenskunstler/banking/domain"
	"github.com/orgs/Lebenskunstler/banking/service"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
