package controller

import (
	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"gorm.io/gorm"
)

type OrderController struct {
	Db *gorm.DB
}

func NewOrderDbController(db *gorm.DB) *OrderController {
	return &OrderController{Db: db}
}
func (o OrderController) CreateOrder(userId, orderId string) error {
	result := o.Db.Create(&domain.Order{
		Id:     orderId,
		UserId: userId,
	})
	return result.Error
}

func (o OrderController) GetOrder(id string) (domain.Order, error) {
	var res domain.Order

	result := o.Db.Where("id= ?", id).Find(&res)
	if result.RowsAffected == 0 {
		return domain.Order{}, result.Error
	}
	return res, nil

}
func (o OrderController) CreateInfoOrder(api domain.FinalResponse) error {
	res, err := o.GetOrder(api.OrderId)
	if err != nil {
		panic(err)
	}
	o.Db.Save(&res)
	result := o.Db.Create(&domain.FinalResponse{
		OrderId:      api.OrderId,
		ProductId:    api.ProductId,
		ActualAmount: api.ActualAmount,
		OrderStatus:  api.OrderStatus,
		PaymentId:    api.PaymentId,
		SenderEmail:  api.SenderEmail,
		OrderTime:    "time.Now()",
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
