package api

import (
	"free5gc-cli/lib/MongoDBLibrary"

	"go.mongodb.org/mongo-driver/bson"
)

func Flush() {
	filter := bson.M{}
	for _, db := range DatabaseCollectionList {
		MongoDBLibrary.RestfulAPIDeleteMany(db, filter)
	}
}

func Drop(db string) {
	filter := bson.M{}
	MongoDBLibrary.RestfulAPIDeleteMany(db, filter)
}
