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
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email      string             `bson:"email" json:"email" binding:"required"`
	Bio        Bio                `bson:"bio" json:"bio" binding:"omitempty"`
	Password   string             `bson:"password" json:"password" binding:"required"`
	Role       string             `bson:"role" json:"role" binding:"required"`
	EmployeeID string             `bson:"employee_id" json:"employee_id" binding:"required"`
}

// Example ... A sample use
var Example = Model{
	ID:       primitive.NewObjectID(),
	Email:    "someone@something.com",
	Password: "password",
}
