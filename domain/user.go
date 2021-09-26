package domain

// model
type User struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"dateofbirth" xml:"dateofbirth"`
	Status      string `json:"status" xml:"status"`
}

type UserRepository interface {
	//secondary port
	GetAllUsers() ([]User, error)
	ById(string) (*User, error)
}
