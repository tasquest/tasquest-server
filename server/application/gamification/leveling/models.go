package leveling

import "github.com/google/uuid"

type ExpLevel struct {
	ID    uuid.UUID `json:"id" bson:"_id"`
	Level int       `json:"level" bson:"level"`
	From  int64     `json:"from" bson:"from"`
	To    int64     `json:"to" bson:"to"`
}
