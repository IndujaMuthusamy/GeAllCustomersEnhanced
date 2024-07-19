package app

import (
	"Banking/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"fullname" xml:"xname"`
	Id   string `json:"identifier" xml:"xid"`
	Age  int    `json:"personage" xml:"xage"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

// class
type CustomerHandlers struct {
	Service service.CustomerService
}

// method associated with customerhandlers class
func (ch CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customerStatus := r.URL.Query().Get("status")
	customers, err := ch.Service.GetAllCustomers(customerStatus)

	/*customer := Customer{
		Name: "INDUJA",
		Id:   "100",
		Age:  20,
	}*/
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	}
	//w.Header().Add("Content-Type", "application/xml")
	//xml.NewEncoder(w).Encode(customer)
	//fmt.Fprint(w, customer)
}

func (ch CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	customer, err := ch.Service.GetCustomerById(customer_id)

	if err != nil {
		log.Println("Encountered Error", err)
		//WriteResponse(w, err.Code, err.Message)

		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprint(w, err.Error())

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err)
		//fmt.Fprint(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
	//w.Header().Add("Content-Type", "application/xml")
	//xml.NewEncoder(w).Encode(customer)
	//fmt.Fprint(w, customer)
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, data)
}
