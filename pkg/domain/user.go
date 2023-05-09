package domain

import (
	"time"
)

type (
	User struct {
		ID           string    `json:"id" gorm:"primary_key"`
		Name         string    `json:"name"`
		Email        string    `json:"email" gorm:"unique"`
		Password     string    `json:"password"`
		RegisteredAt time.Time `json:"registered_at"`
	}
)
