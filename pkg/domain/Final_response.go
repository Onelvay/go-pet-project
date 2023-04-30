package domain

type FinalResponse struct {
	ActualAmount string  `json:"actual_amount"`
	OrderId      string  `json:"order_id"`
	ProductId    string  `json:"product_id"`
	SenderEmail  string  `json:"sender_email"`
	OrderStatus  string  `json:"order_status"`
	Comment      string  `json:"comment"`
	Rating       float32 `json:"rating"`

	Order Order `gorm:"references:Id"`
}
