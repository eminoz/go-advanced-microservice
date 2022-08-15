package service

import (
	"github.com/eminoz/go-redis-project/cache"
	"github.com/eminoz/go-redis-project/model"
	"github.com/eminoz/go-redis-project/repository"
	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	CreateUser(ctx *fiber.Ctx) (interface{}, error)
	GetUserByEmail(ctx *fiber.Ctx) model.UserDal
	GetAllUser(ctx *fiber.Ctx) []model.UserDal
}

type UserService struct {
	UserRepository repository.IUserRepository
	UserRedis      cache.IUserCache
}

func (u *UserService) CreateUser(ctx *fiber.Ctx) (interface{}, error) {
	m := new(model.User)
	ctx.BodyParser(m)
	createUser, err := u.UserRepository.CreateUser(ctx, m)
	if err != nil {
		return nil, err
	}
	dal := model.UserDal{Name: m.Name, Email: m.Email}
	u.UserRedis.SaveUserByEmail(dal)
	return createUser, nil
}
func (u *UserService) GetUserByEmail(ctx *fiber.Ctx) model.UserDal {
	email := ctx.Params("email")
	userByEmail := u.UserRedis.GetUserByEmail(email)
	if userByEmail.Email != "" {
		return userByEmail
	}
	getUserByEmail := u.UserRepository.GetUserByEmail(ctx, email)
	return getUserByEmail
}
func (u UserService) GetAllUser(ctx *fiber.Ctx) []model.UserDal {
	user := u.UserRedis.GetAllUser()
	if len(user) > 0 {
		return user
	}
	getAllUser := u.UserRepository.GetAllUser(ctx)
	return getAllUser
}
