package domain

type Book struct {
	Id          string  `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"desc"`
	// Created_at  time.Time `gorm:"default:current_timestamp"`
}
