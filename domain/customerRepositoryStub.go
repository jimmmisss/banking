package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "0001", Name: "Wesley Pereira", City: "Palhoça/SC", Zipcode: "01", DateOfBirth: "1982-09-13", Status: "1"},
		{Id: "0002", Name: "Isadora Pereira", City: "Pindaré Mirim/MA", Zipcode: "01", DateOfBirth: "2013-05-04", Status: "1"},
		{Id: "0003", Name: "Fadia Pereira", City: "Palhoça/SC", Zipcode: "01", DateOfBirth: "1993-08-05", Status: "1"},
		{Id: "0004", Name: "Isabel Pereira", City: "Santa Inês/MA", Zipcode: "01", DateOfBirth: "1993-08-05", Status: "0"},
	}
	return CustomerRepositoryStub{customers: customers}
}
