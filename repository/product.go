package repository

import (
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type IProductRepository interface {
	CreateProduct()
}

func (c ProductCollection) CreateProduct(ctx fiber.Ctx, p *model.Product) {
	insertOne, err := c.Collection.InsertOne(ctx.Context(), p)
	if err != nil {
	}
	var prod model.Product
	bson.D
	return
}
