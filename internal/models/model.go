package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Tasks       []Task             `json:"tasks" bson:"tasks"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

type Task struct {
	Name     string `json:"name" bson:"name"`
	Priority string `json:"priority" bson:"priority"`
	Duration string `json:"duration" bson:"duration"`
}
