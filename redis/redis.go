package redis

import (
	"github.com/go-redis/redis"
)

func InitRedis(host, pass string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		DB:       0,
	})
	// _, err := client.Ping().Result()
	// if err != nil {
	// 	return client, errors.New("failed to initialize redis")
	// }
	return client, nil
}
