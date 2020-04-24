package adminmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

// Bio : user's biography
type Bio struct {
	firstname  string `bson:"firstname"`
	middlename string `bson:"middlename"`
	lastname   string `bson:"lastname"`
	gender     string `bson:"gender"`
	picture    string `bson:"picture"`
}

// Model : user's model
type Model struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Email      string               `bson:"email" json:"email" binding:"required"`
	Bio        Bio                  `bson:"bio" json:"bio" binding:"omitempty"`
	Password   string               `bson:"password" json:"password" binding:"required"`
	Type       string               `bson:"type" json:"type" binding:"required"`
	Demo       bool                 `bson:"demo" json:"demo" binding:"required"`
	Role       string               `bson:"role" json:"role" binding:"required"`
	Securities []primitive.ObjectID `bson:"securities,omitempty" json:"securities"`
	Trade      []primitive.ObjectID `bson:"trade" json:"trade"`
	Stock      []primitive.ObjectID `bson:"stock" json:"stock"`
}
