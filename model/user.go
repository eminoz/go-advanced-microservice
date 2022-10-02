package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string `json:"name"`
	Email    string `validate:"required,email,omitempty"`
	Password string `validate:"required,gte=7,lte=130,omitempty"`
	Role     string `json:"role"`
	Orders   Orders
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
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Token string             `json:"token"`
}
