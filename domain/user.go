package domain

import "github.com/goddamnnoob/notReddit/errs"

// model
type User struct {
	Id          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"dateofbirth" xml:"dateofbirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type UserRepository interface {
	//secondary port
	GetAllUsers() ([]User, *errs.AppError)
	ById(string) (*User, *errs.AppError)
	ByStatus(int) ([]User, *errs.AppError)
}
