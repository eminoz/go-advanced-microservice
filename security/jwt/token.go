package jwt

import (
	"fmt"
	"github.com/eminoz/go-advanced-microservice/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type IToken interface {
	GenerateJWT(email string, role string) (string, error)
}

//go:generate mockgen -destination=../mocks/Auth/mockUserAuth.go -package=jwt  github.com/eminoz/go-advanced-microservice/security/jwt IToken

type Auth struct{}

func (a Auth) GenerateJWT(email string, role string) (string, error) {
	secretKey := config.GetConfig().AppSecret
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
