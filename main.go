package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Onelvay/docker-compose-project/payment/client"
	contr "github.com/Onelvay/docker-compose-project/pkg/controller"
	handlersss "github.com/Onelvay/docker-compose-project/pkg/handlers"
	"github.com/Onelvay/docker-compose-project/pkg/server"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	db "github.com/Onelvay/docker-compose-project/postgres"
	redisClient "github.com/Onelvay/docker-compose-project/redis"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	config := db.NewConfig(viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
	)
	client.InitConst(viper.GetString("payment.merchantId"), viper.GetString("payment.merchantPassword"), viper.GetString("payment.checkoutUrl"))

	postgres := db.NewPostgresDb(*config)

	redis, err := redisClient.InitRedis(viper.GetString("redis.host"), viper.GetString("redis.password"))
	if err != nil {
		panic(err)
	}

	db, userDb, tokenDb, orderDb := initDbControllers(postgres, redis)
	hasher := service.NewHasher(viper.GetString("app.hash"))

	userContr := contr.NewUserController(userDb, tokenDb, hasher, orderDb)
	handlers := contr.NewHandlers(db, &userContr, orderDb, tokenDb)
	class := handlersss.NewUserHandler(&userContr, userDb)
	router := server.InitRoutes(handlers, *class)
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}
	err = http.ListenAndServe(":"+PORT, router)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDbControllers(postgres *gorm.DB, redis *redis.Client) (*contr.BookstorePostgres, *contr.UserPostgres, *contr.TokenPostgres, *contr.OrderController) {
	db := contr.NewBookstoreDbController(postgres, redis)
	userDb := contr.NewUserDbController(postgres)
	tokenDb := contr.NewTokenDbController(postgres)
	order := contr.NewOrderDbController(postgres)
	return db, userDb, tokenDb, order
}
