package service

import (
	"fmt"
	"github.com/eminoz/go-advanced-microservice/cache"
	"github.com/eminoz/go-advanced-microservice/core/utilities"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/eminoz/go-advanced-microservice/security/encryption"
	"github.com/eminoz/go-advanced-microservice/security/jwt"
	"github.com/gofiber/fiber/v2"
)

//go:generate mockgen -destination=../mocks/services/mockUserService.go -package=service  github.com/eminoz/go-advanced-microservice/service IUserService

type IUserService interface {
	CreateUser(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError)
	GetUserByEmail(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError)
	GetAllUser(ctx *fiber.Ctx) *utilities.ResultOfSuccessData
	DeleteUserByEmail(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError)
	UpdateUserByEmail(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError)
	SignIn(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError)
	GetUsersAddress(ctx *fiber.Ctx) *utilities.ResultOfSuccessData
}

type UserService struct {
	UserRepository repository.IUserRepository
	UserRedis      cache.IUserCache
	Authentication jwt.IToken
	Encryption     encryption.Encryption
}

func NewUserService(r repository.IUserRepository, c cache.IUserCache, j jwt.IToken, e encryption.Encryption) IUserService {
	return &UserService{
		UserRepository: r,
		UserRedis:      c,
		Authentication: j,
		Encryption:     e,
	}
}
func (u UserService) createToken(ctx *fiber.Ctx, email string, password string) (model.Token, *utilities.ResultError) {
	user := u.UserRepository.GetUserByEmailForAuth(ctx, email)
	if user.Email == "" {
		return model.Token{}, utilities.ErrorResult("user not found ")
	}
	checkPasswordHash := u.Encryption.CheckPasswordHash(password, user.Password)
	if !checkPasswordHash {
		return model.Token{}, utilities.ErrorResult("password is incorrect")
	}
	generateJWT, err := u.Authentication.GenerateJWT(user.Email, user.Role)

	if err != nil {
		return model.Token{}, utilities.ErrorResult("did not generate token")
	}
	var token model.Token
	token.Email = user.Email
	token.Role = user.Role
	token.ID = user.ID
	token.TokenString = generateJWT
	return token, nil
}
func (u UserService) SignIn(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError) {
	m := new(model.Authentication)
	ctx.BodyParser(m)
	token, resultError := u.createToken(ctx, m.Email, m.Password)
	if resultError != nil {
		return nil, utilities.ErrorResult(resultError.Message)
	}
	return utilities.SuccessDataResult("signed in successfully", token), nil
}
func (u UserService) CreateUser(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError) {
	m := new(model.User)
	ctx.BodyParser(m)
	email := u.UserRepository.GetUserByEmail(ctx, m.Email)
	if email.Email != "" {
		return nil, utilities.ErrorResult("user already exist")
	}
	password, _ := u.Encryption.GenerateHashPassword(m.Password)
	passwordForAuth := m.Password
	m.Password = password
	createUser, _ := u.UserRepository.CreateUser(ctx, m)
	token, resultError := u.createToken(ctx, m.Email, passwordForAuth)
	if resultError != nil {
		return nil, utilities.ErrorResult(resultError.Message)
	}
	dal := model.UserDal{ID: createUser.ID, Name: m.Name, Email: m.Email, Token: token.TokenString} //model mapping
	u.UserRedis.SaveUserByEmail(dal)
	//save user in redis

	return utilities.SuccessDataResult("user created", dal), nil
}
func (u UserService) GetUsersAddress(ctx *fiber.Ctx) *utilities.ResultOfSuccessData {
	email := ctx.Params("email")
	address := u.UserRepository.GetUserAddress(ctx, email)
	return utilities.SuccessDataResult("address", address)
}
func (u UserService) GetUserByEmail(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError) {
	email := ctx.Params("email")
	userByEmail := u.UserRedis.GetUserByEmail(email)
	if userByEmail.Email == email {
		return utilities.SuccessDataResult("user found", userByEmail), nil
	}
	getUserByEmail := u.UserRepository.GetUserByEmail(ctx, email)
	if getUserByEmail.Email == "" {
		return nil, utilities.ErrorResult("user did not found")
	}
	result := utilities.SuccessDataResult("user found", getUserByEmail)
	return result, nil
}
func (u UserService) GetAllUser(ctx *fiber.Ctx) *utilities.ResultOfSuccessData {
	user := u.UserRedis.GetAllUser()
	if len(user) > 0 {
		result := utilities.SuccessDataResult("all users", user)
		return result
	}
	getAllUser := u.UserRepository.GetAllUser(ctx)
	result := utilities.SuccessDataResult("all users", getAllUser)
	return result
}
func (u UserService) DeleteUserByEmail(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError) {
	email := ctx.Params("email")

	byEmail, _ := u.UserRepository.DeleteUserByEmail(ctx, email)

	if byEmail == 0 {
		return nil, utilities.ErrorResult("user did not find to delete")
	}

	u.UserRedis.DeleteUserByEmail(email) //delete user in redis
	result := utilities.SuccessResult("user deleted")
	return result, nil
}
func (u UserService) UpdateUserByEmail(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError) {
	email := ctx.Params("email")
	m := new(model.UserDal)
	ctx.BodyParser(m)
	if m.Email == "" {
		return nil, utilities.ErrorResult("user mustn't be empty")
	}

	byEmail, msg := u.UserRepository.UpdateUserByEmail(ctx, email, *m)
	if byEmail {
		fmt.Println(m)
		u.UserRedis.SaveUserByEmail(*m) //updated user in redis
		return utilities.SuccessResult(msg), nil
	}

	return nil, utilities.ErrorResult(msg)
}
