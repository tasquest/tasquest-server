package tasks

import (
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Experience  int64     `json:"leveling" bson:"leveling"`
	Active      bool      `json:"active" bson:"active"`
}
