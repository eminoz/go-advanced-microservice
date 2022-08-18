package service

import (
	"fmt"
	"github.com/eminoz/go-advanced-microservice/cache"
	"github.com/eminoz/go-advanced-microservice/core/utilities"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/eminoz/go-advanced-microservice/security/jwt"
	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	CreateUser(ctx *fiber.Ctx) (interface{}, *utilities.ResultError)
	GetUserByEmail(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError)
	GetAllUser(ctx *fiber.Ctx) *utilities.ResultOfSuccessData
	DeleteUserByEmail(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError)
	UpdateUserByEmail(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError)
	SignIn(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError)
}

type UserService struct {
	UserRepository repository.IUserRepository
	UserRedis      cache.IUserCache
	Authentication jwt.IToken
}

func (u UserService) SignIn(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError) {
	m := new(model.Authentication)
	ctx.BodyParser(m)
	user := u.UserRepository.GetUserByEmailForAuth(ctx, m.Email)
	if user.Email == "" {
		return nil, utilities.ErrorResult("user not found ")
	}

	generateJWT, err := u.Authentication.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return nil, utilities.ErrorResult("not generate token")
	}
	var token model.Token
	token.Email = user.Email
	token.Role = user.Role
	token.TokenString = generateJWT
	return utilities.SuccessDataResult("signed in successfully", token), nil
}
func (u *UserService) CreateUser(ctx *fiber.Ctx) (interface{}, *utilities.ResultError) {
	m := new(model.User)
	ctx.BodyParser(m)
	email := u.UserRepository.GetUserByEmail(ctx, m.Email)
	if email.Email != "" {
		return nil, utilities.ErrorResult("user already exist")
	}
	createUser, _ := u.UserRepository.CreateUser(ctx, m)

	go func() {
		dal := model.UserDal{Name: m.Name, Email: m.Email} //model mapping
		u.UserRedis.SaveUserByEmail(dal)                   //save user in redis

	}()

	return utilities.SuccessDataResult("user created", createUser), nil
}
func (u *UserService) GetUserByEmail(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError) {
	email := ctx.Params("email")
	userByEmail := u.UserRedis.GetUserByEmail(email)
	if userByEmail.Email == email {
		fmt.Println(userByEmail)
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

	go func() {
		u.UserRedis.DeleteUserByEmail(email) //delete user in redis
	}()
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
		go func() {
			u.UserRedis.SaveUserByEmail(*m) //updated user in redis
		}()
		return utilities.SuccessResult(msg), nil
	}

	return nil, utilities.ErrorResult(msg)
}
