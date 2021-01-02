package profiles

import "time"

type CreateUserProfile struct {
	UserID   string    `json:"user_id" validate:"required"`
	Name     string    `json:"name" validate:"required,alpha"`
	Surname  string    `json:"surname" validate:"required,alpha"`
	Birthday time.Time `json:"birthday" validate:"required,datetime"`
}
