package repository

import (
	"github.com/eminoz/go-advanced-microservice/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func UserCollectionSetting() *UserCollection {
	getDatabase := database.GetDatabase()
	return &UserCollection{
		Db:         getDatabase,
		Collection: getDatabase.Collection("user"),
	}
}
