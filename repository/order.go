package repository

/**/
import (
	"context"
	"fmt"
	"github.com/eminoz/go-advanced-microservice/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrders interface {
	CreateNewOrdersById(ctx fiber.Ctx, id string, update interface{}) (interface{}, error)
	GetUsersOrders(ctx context.Context, id string) model.Orders
}

func (u UserCollection) CreateNewOrdersById(ctx fiber.Ctx, id string, update interface{}) (interface{}, error) {
	userID, _ := primitive.ObjectIDFromHex(id)
	d := bson.D{{"_id", userID}}
	b := bson.D{{"$set", bson.D{{"orders", update}}}}
	updateOne, err := u.Collection.UpdateOne(ctx.Context(), d, b)
	if err != nil {
		return nil, err
	}
	return updateOne, nil

}
func (u UserCollection) GetUsersOrders(ctx context.Context, id string) model.Orders {
	userID, _ := primitive.ObjectIDFromHex(id)
	d := bson.D{{"_id", userID}}
	var user model.User

	err := u.Collection.FindOne(ctx, &d).Decode(&user)
	fmt.Println(user.Orders)
	if err != nil {
		fmt.Println(err)
	}
	return user.Orders
}
