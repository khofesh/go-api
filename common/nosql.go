package common

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MoDB : reference to mongodb client
var MoDB *mongo.Client

// RedisClient ...
var RedisClient *redis.Client

// InitMongo : initiate connection to mongodb
func InitMongo(mongoURI string) error {
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return errors.New("Error connecting to mongodb")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return errors.New("Pinging to mongodb error")
	}
	fmt.Println("Connected to MongoDB!")

	MoDB = client

	return nil
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

// InitRedis : initialize connection to redis
func InitRedis() error {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

// GetRedis ...
func GetRedis() *redis.Client {
	return RedisClient
}
