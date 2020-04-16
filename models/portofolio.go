package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PortofolioModel ...
type PortofolioModel struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	BuyDate        time.Time          `bson:"buydate" json:"buydate"`
	Shares         float64            `bson:"shares" json:"shares"`
	BuyPrice       float64            `bson:"buyprice" json:"buyprice"`
	AvePrice       float64            `bson:"aveprice" json:"aveprice"`
	BuyComm        float64            `bson:"buycomm" json:"buycomm"`
	SellPrice      float64            `bson:"sellprice" json:"sellprice"`
	SellComm       float64            `bson:"sellcomm" json:"sellcomm"`
	Profit         float64            `bson:"profit" json:"profit"`
	ROI            float32            `bson:"roi" json:"roi"`
	BreakSellPrice float64            `bson:"breaksellprice" json:"breaksellprice"`
	SellDate       time.Time          `bson:"selldate" json:"selldate"`
}
