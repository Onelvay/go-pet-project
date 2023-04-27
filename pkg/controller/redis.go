package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"github.com/go-redis/redis"
)

func productExistInRedis(client *redis.Client, id uint) (domain.Product, error) {
	var product domain.Product
	val, err := client.Get(fmt.Sprint(id)).Result()
	if err == nil {
		err = json.Unmarshal([]byte(val), &product)
		if err == nil {
			return product, nil
		}
		return product, err
	}
	return product, err
}
func saveBookInRedis(r *redis.Client, product domain.Product) {
	j, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
	}
	err = r.Set(fmt.Sprint(product.Id), j, 0).Err()
	if err != nil {
		log.Println(err)
	}
}
