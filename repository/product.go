package repository

import (
	"fmt"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type IProductRepository interface {
	CreateProduct(ctx *fiber.Ctx, p *model.Product) model.Product
	GetAllProduct(ctx *fiber.Ctx) []model.ProductDal
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

func (c ProductCollection) GetAllProduct(ctx *fiber.Ctx) []model.ProductDal {
	find, err := c.Collection.Find(ctx.Context(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	var p []model.ProductDal
	err2 := find.All(ctx.Context(), &p)
	fmt.Println(p)
	if err2 != nil {
		fmt.Println(err2)
	}
	return p
}
