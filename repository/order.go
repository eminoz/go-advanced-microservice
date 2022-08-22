package repository

/**/
import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrders interface {
	CreateNewOrdersById(ctx fiber.Ctx, id string, update interface{}) (interface{}, error)
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
