package domain

type FinalResponse struct {
	ActualAmount string `json:"actual_amount"`
	OrderId      string `json:"order_id"`
	PaymentId    int    `json:"payment_id"`
	ProductId    string `json:"product_id"`
	SenderEmail  string `json:"sender_email"`
	OrderStatus  string `json:"order_status"`
	OrderTime    string `json:"order_time"`

	Order Order `gorm:"references:Id"`
}
