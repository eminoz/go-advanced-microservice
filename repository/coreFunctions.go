package repository

import "go.mongodb.org/mongo-driver/bson"

func makeFilter[filter any](field string, f filter) bson.D {
	d := bson.D{{field, f}}
	return d
}

func makeFilterAndUpdate[filter any, update any](field string, e string, f filter, u update) (bson.D, bson.D) {
	d := bson.D{{field, f}}
	up := bson.D{{e, u}}
	return d, up

}
