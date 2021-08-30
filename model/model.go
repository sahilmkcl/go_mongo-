package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name"`
	LastName string             `json:"lastName"`
	Password string             `json:"password"`
}

type Update struct {
	ToUpdate string `json:"toUpdate"`
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}
