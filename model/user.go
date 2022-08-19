package model

type User struct {
	Name     string `json:"name"`
	Email    string `validate:"required,email,omitempty"`
	Password string `validate:"required,gte=7,lte=130,omitempty"`
	Role     string `json:"role"`
}
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
type UserDal struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
