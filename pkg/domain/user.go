package domain

import (
	"time"
)

type (
	User struct {
		ID           string    `json:"id"`
		Name         string    `json:"name"`
		Email        string    `json:"email"`
		Password     string    `json:"password"`
		RegisteredAt time.Time `json:"registered_at"`
	}

	UserOrders struct {
		ProductId string
	}
)
