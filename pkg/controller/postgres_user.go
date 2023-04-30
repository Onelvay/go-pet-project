package controller

import (
	"context"
	"errors"
	"strings"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserDbController(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(cnt context.Context, user domain.User) error {
	byteid := uuid.New()
	id := strings.Replace(byteid.String(), "-", "", -1)
	res := r.db.Create(&domain.User{
		ID:           id,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		RegisteredAt: user.RegisteredAt,
	})
	return res.Error
}
func (r *UserPostgres) SignInUser(cxt context.Context, email, password string) (domain.User, error) {
	var user domain.User
	res := r.db.Where("email = ? AND password = ?", email, password).Find(&user)
	return user, res.Error
}

func (r *UserPostgres) GetUserOrders(id string) ([]uint, error) {
	var orders []uint
	r.db.Table("final_responses").Select("final_responses.product_id").Joins("INNER JOIN orders ON orders.id = final_responses.order_id").Where("orders.user_id = ?", id).Scan(&orders)
	return orders, nil
}

func (o *UserPostgres) AddDetailToOrder(req domain.OrderDetail) error {
	var res domain.Order
	result := o.db.Where("id = ? and user_id = ?", req.Order_id, req.User_id).Find(&res)
	if result.RowsAffected == 1 {
		if err := o.db.Model(&domain.FinalResponse{}).Where("order_id = ?", req.Order_id).Update("comment", req.Comment).Error; err != nil {
			return err
		}
		if err := o.db.Model(&domain.FinalResponse{}).Where("order_id = ?", req.Order_id).Update("rating", req.Rating).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("y пользователя нет такого заказа")
}
