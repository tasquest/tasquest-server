package tasks

import (
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Experience  int64     `json:"experience" bson:"experience"`
}

type ExperienceTable struct {
	ID    uuid.UUID `json:"id" bson:"_id"`
	Level int       `json:"level" bson:"level"`
	From  int64     `json:"from" bson:"from"`
	To    int64     `json:"to" bson:"to"`
}
