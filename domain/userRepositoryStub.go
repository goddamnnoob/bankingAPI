package domain

// stub is placeholder that simulates the input
type UserRepositoryStub struct {
	users []User
}

func (u UserRepositoryStub) GetAllUsers() ([]User, error) {
	return u.users, nil
}

func NewUserRepositoryStub() UserRepositoryStub {
	users := []User{
		{"Gowtham", "1", "mrmu", "gowtham@me.com"},
		{"ZOOM", "2", "zoom", "zoom@zoom.com"},
	}
	return UserRepositoryStub{users}
}
