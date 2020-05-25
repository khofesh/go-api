package stmodels

import "go.mongodb.org/mongo-driver/bson/primitive"

// Securities ...
type Securities struct {
	ID       primitive.ObjectID   `bson:"_id" json:"id"`
	Name     string               `bson:"name" json:"name"`
	BuyCost  float32              `bson:"buycost" json:"buycost"`
	SellCost float32              `bson:"sellcost" json:"sellcost"`
	Stocks   []primitive.ObjectID `bson:"stocks" json:"stocks"`
}
