package api

import (
	"github.com/eminoz/go-advanced-microservice/service"
	"github.com/gofiber/fiber/v2"
)

type IOrderController interface {
	CreateOrder(ctx *fiber.Ctx) error
}
type OrderController struct {
	OrderService service.IOrderService
}

func NewOrderController(s service.IOrderService) IOrderController {
	return &OrderController{
		OrderService: s,
	}

}

func (o OrderController) CreateOrder(ctx *fiber.Ctx) error {
	id := o.OrderService.CreateNewOrdersById(*ctx)
	return ctx.JSON(id)
}
