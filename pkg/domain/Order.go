package domain

type Order struct {
	OrderId     string `gorm:"primary_key" json:"orderId"`
	UserId      string `json:"userId"`
	OrderStatus string `json:"order_status"`
	User        User   `gorm:"references:ID"`
}
