package controller

import (
	"context"
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

func (r *UserPostgres) GetUserOrders(id string) ([]domain.UserOrders, error) {
	var orders []domain.UserOrders
	// rows, err := r.db.Table("final_responses").Select("final_responses.product_id").Joins("join on orders.id=final_responses.order_id AND orders.user_id = ?", id).Rows()
	r.db.Table("final_responses").Select("final_responses.product_id").Joins("INNER JOIN orders ON orders.id = final_responses.order_id").Where("orders.user_id = ?", id).Scan(&orders)
	// r.db.InnerJoins("orders").Find(&orders)
	// if err != nil {
	// 	return []domain.FinalResponse{}, err
	// }
	// err = rows.Scan(&orders)
	// if err != nil {
	// 	return []domain.FinalResponse{}, err
	// }
	// fmt.Println(orders)
	return orders, nil
}
