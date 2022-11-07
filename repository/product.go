package repository

import (
	"fmt"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	CreateProduct(ctx *fiber.Ctx, p *model.Product) model.Product
	GetAllProduct(ctx *fiber.Ctx) []model.ProductDal
	UpdateProductBProductName(ctx *fiber.Ctx, email string, p *model.Product) *mongo.UpdateResult
	DeleteProduct(ctx *fiber.Ctx, productname string) bool
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
func (c ProductCollection) UpdateProductBProductName(ctx *fiber.Ctx, productname string, p *model.Product) *mongo.UpdateResult {
	filter := bson.D{{"productname", productname}}
	update := bson.D{{"$set", p}}
	one, err := c.Collection.UpdateOne(ctx.Context(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
	return one

}
func (c ProductCollection) GetAllProduct(ctx *fiber.Ctx) []model.ProductDal {
	find, err := c.Collection.Find(ctx.Context(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	var p []model.ProductDal
	err2 := find.All(ctx.Context(), &p)
	if err2 != nil {
		fmt.Println(err2)
	}
	return p
}
func (c ProductCollection) DeleteProduct(ctx *fiber.Ctx, productname string) bool {
	filter := bson.D{{"productname", productname}}
	one, _ := c.Collection.DeleteOne(ctx.Context(), filter)
	if one.DeletedCount == 1 {
		return true
	}
	return false
}
