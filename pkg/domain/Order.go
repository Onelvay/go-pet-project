package domain

type Order struct {
	Id     string `gorm:"primary_key" json:"orderId"`
	UserId string `json:"userId"`
	User   User   `gorm:"references:ID"`
}
