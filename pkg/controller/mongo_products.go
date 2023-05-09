package controller

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"

	// mongoDb "github.com/Onelvay/docker-compose-project/db/mongoDB"
	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
)

type ProductDBController struct {
	mongo       *mongo.Collection
	redisClient *redis.Client
	postgres    *gorm.DB
}

func NewProductDbController(collection *mongo.Collection, redis *redis.Client, postgres *gorm.DB) *ProductDBController {
	return &ProductDBController{collection, redis, postgres}
}

func (p *ProductDBController) GetProductById(id uint64) (domain.Product, error) {
	product, err := getProductFromRedisById(p.redisClient, id)
	if err == nil {
		return product, nil
	}

	if err = p.mongo.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product); err != nil {
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
	cursor, err := p.mongo.Find(context.Background(), bson.M{"name": name})
	if err != nil {
		return products, err
	}
	if err = cursor.All(context.Background(), &products); err != nil {
		return products, err
	}
	saveProductsInRedis(p.redisClient, name, products)
	return products, nil
}

func (p *ProductDBController) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	cur, err := p.mongo.Find(context.Background(), bson.D{})
	if err != nil {
		return products, err
	}
	if err = cur.All(context.Background(), &products); err != nil {
		return products, err
	}
	return products, nil
}
func (p *ProductDBController) CreateProduct(product domain.Product) error {
	_, err := p.mongo.InsertOne(context.Background(), product)
	if err != nil {
		return err
	}
	saveProductInRedis(p.redisClient, product)
	return err
}
func (p *ProductDBController) DeleteProductById(id uint64) error {
	_, err := p.mongo.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
func (p *ProductDBController) GetProductRating(id uint) float64 {
	var result float64
	p.postgres.Table("final_responses").Where("product_id = ?", fmt.Sprint(id)).Select("AVG(rating)").Group("product_id").Pluck("AVG(rating)", &result)

	return result
}
