package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/go-redis/redis"
)

func getProductFromRedisById(client *redis.Client, id uint64) (domain.Product, error) {
	var product domain.Product
	val, err := client.Get(fmt.Sprint(id)).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &product); err == nil {
			return product, nil
		}
		return product, err
	}
	return product, err
}
func getProductsFromRedisByName(client *redis.Client, name string) ([]domain.Product, error) {
	var products []domain.Product
	val, err := client.Get(name).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &products); err == nil {
			return products, nil
		}
		return products, err
	}
	return products, err
}
func saveProductInRedis(r *redis.Client, product domain.Product) {
	j, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
	}
	err = r.Set(fmt.Sprint(product.Id), j, 0).Err()
	if err != nil {
		log.Println(err)
	}
}
func saveProductsInRedis(r *redis.Client, name string, products []domain.Product) {
	j, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
	}
	err = r.Set(name, j, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}
