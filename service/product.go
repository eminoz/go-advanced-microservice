package service

import (
	"github.com/eminoz/go-advanced-microservice/core/utilities"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/gofiber/fiber/v2"
)

type IProductService interface {
	CreateProduct(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultOfErrorData)
}
type ProductService struct {
	ProductRepository repository.IProductRepository
}

func NewProductService(p repository.IProductRepository) IProductService {
	return &ProductService{ProductRepository: p}
}
func (s ProductService) CreateProduct(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultOfErrorData) {
	m := new(model.Product)
	err := ctx.BodyParser(&m)
	if err != nil {
		return nil, utilities.ErrorDataResult("some got wrong", err)
	}
	product := s.ProductRepository.CreateProduct(ctx, m)
	dal := model.ProductDal{ProductName: product.ProductName, Quantity: product.Quantity, Price: product.Price, Description: product.Description}

	return utilities.SuccessDataResult("Product Created", dal), nil
}
