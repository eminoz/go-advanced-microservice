package service

import (
	"context"
	"fmt"
	"github.com/eminoz/go-advanced-microservice/core/utilities"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/gofiber/fiber/v2"
)

type IOrderService interface {
	CreateNewOrdersById(ctx fiber.Ctx) interface{}
	GetOrders(ctx fiber.Ctx) *utilities.ResultOfSuccessData
}

type OderService struct {
	OrderRepository repository.IOrders
}

func NewOrderService(o repository.IOrders) IOrderService {
	return &OderService{
		OrderRepository: o,
	}
}

func (o OderService) CreateNewOrdersById(ctx fiber.Ctx) interface{} {
	userID := ctx.Params("id")
	m := new(model.Orders)
	ctx.BodyParser(m)
	fmt.Println(m)
	ordersById, err := o.OrderRepository.CreateNewOrdersById(ctx, userID, m)
	if err != nil {
		result := utilities.ErrorResult("did not updated")
		return result
	}
	return utilities.SuccessDataResult("user updated", ordersById)
}
func (o OderService) GetOrders(ctx fiber.Ctx) *utilities.ResultOfSuccessData {
	context := context.TODO()
	userId := ctx.Params("id")
	orders := o.OrderRepository.GetUsersOrders(context, userId)
	return utilities.SuccessDataResult("user orders", orders)

}
