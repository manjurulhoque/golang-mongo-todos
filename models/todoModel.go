package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     *string            `json:"title" validate:"required,min=2,max=140"`
	Content   *string            `json:"content" validate:"required,min=2"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
