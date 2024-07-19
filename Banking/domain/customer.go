package domain

import "Banking/customererrs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}
type CustomerRepo interface {
	FindAll(status string) ([]Customer, *customererrs.AppError)
	FindById(customer_id string) (*Customer, *customererrs.AppError)
}
