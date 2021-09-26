package domain

//mock adapter
// stub is placeholder that simulates the input
type UserRepositoryStub struct {
	users []User
}

func (u UserRepositoryStub) GetAllUsers() ([]User, error) {
	return u.users, nil
}

func NewUserRepositoryStub() UserRepositoryStub {
	users := []User{
		{"1", "Gowtham", "chennai", "33333", "2.2.2002", "available"},
		{"2", "Gotham", "chennai", "33332", "2.2.2003", "available"},
	}
	return UserRepositoryStub{users}
}
