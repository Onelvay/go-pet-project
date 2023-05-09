package main

import (
	"fmt"
	routes "github.com/Onelvay/docker-compose-project/pkg/http/routes"
	"log"
	"net/http"
	"os"

	mongoDb "github.com/Onelvay/docker-compose-project/db/mongoDB"
	"github.com/Onelvay/docker-compose-project/db/postgres"
	"github.com/Onelvay/docker-compose-project/payment/client"
	contr "github.com/Onelvay/docker-compose-project/pkg/controller"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	redisClient "github.com/Onelvay/docker-compose-project/redis"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	config := postgres.NewConfig(viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
	)
	client.InitConst(viper.GetString("payment.merchantId"), viper.GetString("payment.merchantPassword"), viper.GetString("payment.checkoutUrl"))

	mongoProductDb := mongoDb.MongoProductCollection(viper.GetString("mongoDB.host"))
	postgresDb := postgres.NewPostgresDb(*config)

	redis, err := redisClient.InitRedis(viper.GetString("redis.host"), viper.GetString("redis.password"))
	if err != nil {
		panic(err)
	}

	productDb, userDb, tokenDb, orderDb := initDbControllers(postgresDb, redis, mongoProductDb)
	hasher := service.NewHasher(viper.GetString("app.hash"))
	userContr := contr.NewUserController(userDb, tokenDb, hasher, orderDb)
	handlers := contr.NewHandlers(productDb, &userContr, orderDb, tokenDb, userDb)
	router := routes.InitRoutes(handlers)

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

func initDbControllers(postgres *gorm.DB, redis *redis.Client, mongo *mongo.Collection) (*contr.ProductDBController, *contr.UserPostgres, *contr.TokenPostgres, *contr.OrderController) {
	productDb := contr.NewProductDbController(mongo, redis, postgres)
	userDb := contr.NewUserDbController(postgres)
	tokenDb := contr.NewTokenDbController(postgres)
	order := contr.NewOrderDbController(postgres)
	return productDb, userDb, tokenDb, order
}
