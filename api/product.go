package api

import (
	"github.com/eminoz/go-advanced-microservice/service"
	"github.com/gofiber/fiber/v2"
)

type IProductController interface {
	CreateProduct(ctx *fiber.Ctx) error
	GetAllProduct(ctx *fiber.Ctx) error
}
type ProductController struct {
	ProductService service.IProductService
}

func NewProductController(p service.IProductService) IProductController {
	return &ProductController{ProductService: p}
}

func (p ProductController) CreateProduct(ctx *fiber.Ctx) error {

	createProduct, errorData := p.ProductService.CreateProduct(ctx)
	if errorData != nil {
		return ctx.JSON(errorData)
	}
	return ctx.JSON(createProduct)

}
func (p ProductController) GetAllProduct(ctx *fiber.Ctx) error {
	product := p.ProductService.GetAllProduct(ctx)
	return ctx.JSON(product)
}
