package controller

import (
	request "github.com/Onelvay/docker-compose-project/payment/APIrequest"
	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"gorm.io/gorm"
)

type OrderController struct {
	Db *gorm.DB
}

func NewOrderDbController(db *gorm.DB) *OrderController {
	return &OrderController{Db: db}
}
func (o OrderController) CreateOrder(userId, orderId string) {
	o.Db.Create(&domain.Order{
		Id:     orderId,
		UserId: userId,
	})
}

func (o OrderController) GetOrder(id string) domain.Order {
	var res domain.Order

	o.Db.Where("id= ?", id).Find(&res)
	return res

}
func (o OrderController) CreateInfoOrder(api request.FinalResponse) {
	res := o.GetOrder(api.OrderId)

	o.Db.Save(&res)
	o.Db.Create(&request.FinalResponse{
		OrderId:      api.OrderId,
		ProductId:    api.ProductId,
		ActualAmount: api.ActualAmount,
		OrderStatus:  api.OrderStatus,
		PaymentId:    api.PaymentId,
		SenderEmail:  api.SenderEmail,
		OrderTime:    "time.Now()",
	})
}
