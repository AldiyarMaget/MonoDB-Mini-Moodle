package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
}
