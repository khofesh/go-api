package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// SecuritiesModel ...
type SecuritiesModel struct {
	ID          primitive.ObjectID   `bson:"_id" json:"id"`
	Name        string               `bson:"name" json:"name"`
	BuyCost     float32              `bson:"buycost" json:"buycost"`
	SellCost    float32              `bson:"sellcost" json:"sellcost"`
	Portofolios []primitive.ObjectID `bson:"portofolios" json:"portofolios"`
}
