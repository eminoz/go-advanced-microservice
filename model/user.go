package model

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserDal struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
