package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
