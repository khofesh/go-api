package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/khofesh/simple-go-api/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// StockModel ... salah nama, should've been stock
type StockModel struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	EntityName   string             `bson:"entityName" json:"entityName"`
	EntityCode   string             `bson:"entityCode" json:"entityCode"`
	EntityNumber string             `bson:"entityNumber" json:"entityNumber"`
	Sector       string             `bson:"sector" json:"sector"`
	Subsector    string             `bson:"subsector" json:"subsector"`
}

// CreateStock ...
func (stock *StockModel) CreateStock() error {
	coll := common.GetCollection("simple", "stock")

	if val, _ := coll.CountDocuments(context.TODO(), bson.M{"entityName": stock.EntityName}); val != 0 {
		return errors.New("This stock already exists")
	}

	idxMod := []mongo.IndexModel{
		{Keys: bson.M{"entityName": 1}, Options: options.Index().SetUnique(true)},
		{Keys: bson.M{"entityCode": 1}, Options: options.Index().SetUnique(true)},
		{Keys: bson.M{"entityNumber": 1}, Options: options.Index().SetUnique(true)},
	}

	names, err := coll.Indexes().CreateMany(context.TODO(), idxMod)
	if err != nil {
		return err
	}
	fmt.Printf("created indexes %v\n", names)

	_, err = coll.InsertOne(context.TODO(), stock)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStock ...
func (stock *StockModel) UpdateStock(update bson.M) error {
	coll := common.GetCollection("simple", "stock")

	_, err := coll.UpdateOne(context.TODO(), bson.M{"entityName": stock.EntityName}, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteStock ...
func (stock *StockModel) DeleteStock() error {
	coll := common.GetCollection("simple", "stock")

	_, err := coll.DeleteOne(context.TODO(), bson.M{"entityName": stock.EntityName})
	if err != nil {
		return err
	}

	return nil
}
