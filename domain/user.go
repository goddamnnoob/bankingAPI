package domain

import (
	"github.com/goddamnnoob/notReddit/dto"
	"github.com/goddamnnoob/notReddit/errs"
)

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

func (u User) toStatusAsText() string {
	statusAsText := "active"
	if u.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (u User) ToDto() dto.UserResponse {
	statusAsText := u.toStatusAsText()
	return dto.UserResponse{
		Id:          u.Id,
		Name:        u.Name,
		City:        u.City,
		Zipcode:     u.Zipcode,
		DateOfBirth: u.DateOfBirth,
		Status:      statusAsText,
	}
}
