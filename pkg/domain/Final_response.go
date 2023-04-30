package domain

import "time"

type FinalResponse struct {
	ActualAmount string    `json:"actual_amount"`
	OrderId      string    `json:"order_id"`
	PaymentId    string    `json:"payment_id"`
	ProductId    string    `json:"product_id"`
	SenderEmail  string    `json:"sender_email"`
	OrderStatus  string    `json:"order_status"`
	OrderTime    time.Time `json:"order_time"`
	Comment      string    `json:"comment"`
	Rating       float32   `json:"rating"`

	Order Order `gorm:"references:Id"`
}
