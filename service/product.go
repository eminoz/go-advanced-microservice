package service

import (
	"github.com/eminoz/go-advanced-microservice/core/utilities"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/eminoz/go-advanced-microservice/repository"
	"github.com/gofiber/fiber/v2"
)

type IProductService interface {
	CreateProduct(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultOfErrorData)
	GetAllProduct(ctx *fiber.Ctx) *utilities.ResultOfSuccessData
	UpdateProductBProductName(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError)
	DeleteProduct(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError)
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
func (s ProductService) UpdateProductBProductName(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError) {
	email := ctx.Params("productname")
	m := new(model.Product)
	ctx.BodyParser(&m)
	updateProductByEmail := s.ProductRepository.UpdateProductBProductName(ctx, email, m)

	if updateProductByEmail.ModifiedCount == 1 {
		return utilities.SuccessResult("product updated"), nil
	}
	return nil, utilities.ErrorResult("product did not update ")
}
func (s ProductService) GetAllProduct(ctx *fiber.Ctx) *utilities.ResultOfSuccessData {
	product := s.ProductRepository.GetAllProduct(ctx)
	return utilities.SuccessDataResult("all product", product)

}

func (s ProductService) DeleteProduct(ctx *fiber.Ctx) (*utilities.ResultSuccess, *utilities.ResultError) {
	productname := ctx.Params("productname")
	product := s.ProductRepository.DeleteProduct(ctx, productname)
	if product == true {
		return utilities.SuccessResult("product deleted"), nil
	}
	return nil, utilities.ErrorResult("product did not deleted")
}
