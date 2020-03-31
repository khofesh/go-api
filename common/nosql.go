package common

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MoDB : reference to mongodb client
var MoDB *mongo.Client

// InitMongo : initiate connection to mongodb
func InitMongo(mongoURI string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	MoDB = client
	return MoDB
}

// GetMongoDB :
// get a connection to mongodb
func GetMongoDB() *mongo.Client {
	return MoDB
}

// GetCollection :
// get collection
func GetCollection(dbname string, collection string) *mongo.Collection {
	return MoDB.Database(dbname).Collection(collection)
}
