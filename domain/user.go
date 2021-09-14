package domain

// model
type User struct {
	Name     string `json:"name" xml:"name"`
	UID      string `json:"uid" xml:"uid"`
	Username string `json:"username" xml:"username"`
	Email    string `json:"email" xml:"email"`
}
