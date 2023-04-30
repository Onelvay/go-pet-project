package domain

type (
	Order struct {
		Id     string `gorm:"primary_key" json:"orderId"`
		UserId string `json:"userId"`
		User   User   `gorm:"references:ID"`
	}
	OrderDetail struct {
		Order_id string  `json:"order_id"`
		Rating   float32 `json:"rating"`
		Comment  string  `json:"comment"`
		User_id  string
	}
)
