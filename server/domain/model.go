package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty"bsonL"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}
