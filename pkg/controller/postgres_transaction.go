package controller

import (
	"time"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"gorm.io/gorm"
)

type OrderController struct {
	Db *gorm.DB
}

func NewOrderDbController(db *gorm.DB) *OrderController {
	return &OrderController{Db: db}
}
func (o *OrderController) CreateOrder(userId, orderId string) error {
	result := o.Db.Create(&domain.Order{
		Id:     orderId,
		UserId: userId,
	})
	return result.Error
}

func (o *OrderController) GetOrder(id string) (domain.Order, error) {
	var res domain.Order
	result := o.Db.Where("id= ?", id).Find(&res)
	return res, result.Error
}
func (o *OrderController) CreateInfoOrder(api domain.FinalResponse) error {
	result := o.Db.Create(&domain.FinalResponse{
		OrderId:      api.OrderId,
		ProductId:    api.ProductId,
		ActualAmount: api.ActualAmount,
		OrderStatus:  api.OrderStatus,
		PaymentId:    api.PaymentId,
		SenderEmail:  api.SenderEmail,
		OrderTime:    time.Now(),
	})
	return result.Error
}
