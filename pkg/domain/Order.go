package domain

type Order struct {
	OrderId   string `gorm:"primary_key" json:"orderId"`
	ProductId string `json:"productId"`
	Buyer     string `json:"buyerId"`
}
