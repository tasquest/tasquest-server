package leveling

import (
	"time"

	"github.com/google/uuid"
)

type ExpLevel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Level     int       `json:"level"`
	StartExp  int64     `json:"start_exp"`
	EndExp    int64     `json:"end_exp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
