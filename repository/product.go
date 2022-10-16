package repository

import (
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type IProductRepository interface {
	CreateProduct(ctx *fiber.Ctx, p *model.Product) model.Product
}

func (c ProductCollection) CreateProduct(ctx *fiber.Ctx, p *model.Product) model.Product {
	insertOne, err := c.Collection.InsertOne(ctx.Context(), p)
	if err != nil {
	}
	var prod model.Product
	d := bson.D{{"_id", insertOne.InsertedID}}
	c.Collection.FindOne(ctx.Context(), d).Decode(&prod)
	return prod
}
