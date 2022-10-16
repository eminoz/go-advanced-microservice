package repository

import (
	"github.com/eminoz/go-advanced-microservice/database"
	"go.mongodb.org/mongo-driver/mongo"
)

//to generate this mock run => go generate ./...
//go:generate mockgen -destination=../mocks/repository/mockUserRepository.go -package=repository  github.com/eminoz/go-advanced-microservice/repository IUserRepository

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

type ProductCollection struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func ProductCollectionSetting() *ProductCollection {
	getDatabase := database.GetDatabase()
	return &ProductCollection{
		Db:         getDatabase,
		Collection: getDatabase.Collection("product"),
	}
}
