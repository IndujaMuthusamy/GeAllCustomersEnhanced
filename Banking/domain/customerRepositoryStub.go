package domain

//repo CLASS
type CustomerRepositoryStub struct {
	customers []Customer
}

//CREATE OBJECT OF repoCLASS used in app.go injection //constructor injection repo class (nodb hence manually creating object)
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customersArr := []Customer{
		{Id: "100", Name: "induja", City: "chennai"},
		{Id: "101", Name: "san", City: "trc"},
		{Id: "102", Name: "chett", City: "bang"},
	}
	return CustomerRepositoryStub{
		customers: customersArr,
	}
}

//CustomerRepositoryStub implements customerrepo
func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}
