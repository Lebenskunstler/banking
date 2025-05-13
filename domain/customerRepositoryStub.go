package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"152", "Alex", "Casa", "1555", "10-10-2010", "1"},
		{"152", "Karl", "Blanca", "1555", "10-10-2010", "1"},
	}
	return CustomerRepositoryStub{customers}
}
