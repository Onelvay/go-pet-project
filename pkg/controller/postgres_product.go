package controller

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	// mongoDb "github.com/Onelvay/docker-compose-project/db/mongoDB"
	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
)

type ProductDBController struct {
	db          *mongo.Collection
	redisClient *redis.Client
}

func NewProductDbController(collection *mongo.Collection, redis *redis.Client) *ProductDBController {
	return &ProductDBController{collection, redis}
}

func (p *ProductDBController) GetProductById(id uint64) (domain.Product, error) {
	product, err := getProductFromRedisById(p.redisClient, id)
	if err == nil {
		return product, nil
	}

	if err = p.db.FindOne(context.Background(), bson.M{"id": id}).Decode(&product); err != nil {
		return product, errors.New(fmt.Sprint("no product with this id :", id))
	}

	saveProductInRedis(p.redisClient, product)
	return product, nil
}

func (p *ProductDBController) GetProductsByName(name string) ([]domain.Product, error) {
	rproducts, err := getProductsFromRedisByName(p.redisClient, name)
	if err == nil {
		return rproducts, nil
	}
	var products []domain.Product
	cursor, err := p.db.Find(context.Background(), bson.M{"name": name})
	if err != nil {
		return products, err
	}
	if err = cursor.All(context.Background(), &products); err != nil {
		return products, err
	}
	saveProductsInRedis(p.redisClient, name, products)
	return products, nil
}

func (p *ProductDBController) GetProducts(sorted bool) ([]domain.Product, error) {
	var products []domain.Product
	cur, err := p.db.Find(context.Background(), bson.D{})
	if err != nil {
		return products, err
	}
	if err = cur.All(context.Background(), &products); err != nil {
		return products, err
	}
	return products, nil
}
func (p *ProductDBController) CreateProduct(product domain.Product) error {
	_, err := p.db.InsertOne(context.Background(), product)
	if err != nil {
		fmt.Println(err)
		return err
	}
	saveProductInRedis(p.redisClient, product)
	return err
}

// func (r *BookstorePostgres) DeleteBookById(id string) error {
// 	res := r.Db.Where("id=?", id).Delete(&domain.Book{})
// 	return res.Error
// }

// func (r *BookstorePostgres) UpdateBook(id string, name string, desc string, price float64) error {
// 	_, res := r.GetBookById(id)
// 	if res == nil {
// 		if name != "" {
// 			book.Name = name
// 		}
// 		if desc != "" {
// 			book.Description = desc
// 		}
// 		if price != 0 {
// 			book.Price = price
// 		}
// 		res := r.Db.Save(&book)
// 		saveBookInRedis(r.redisClient)
// 		return res.Error
// 	}
// 	return res
// }
