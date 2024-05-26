package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID     primitive.ObjectID `json: "_id,omitempty" bson:"_id,omitempty" `
	task   string             `json: "task,omitempty" `
	status bool               `json: "status,omitempty" `
}
