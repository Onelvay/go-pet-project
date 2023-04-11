package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Onelvay/docker-compose-project/payment/client"
	"github.com/Onelvay/docker-compose-project/pkg/server"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	"github.com/spf13/viper"

	contr "github.com/Onelvay/docker-compose-project/pkg/controller"
	db "github.com/Onelvay/docker-compose-project/postgres"
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

	db := contr.NewBookstoreDbController(postgres)
	userDb := contr.NewUserDbController(postgres)
	tokenDb := contr.NewTokenDbController(postgres)
	hasher := contr.NewHasher(viper.GetString("app.hash"))
	order := contr.NewOrderDbController(postgres)

	userContr := service.NewUserController(userDb, tokenDb, hasher, order)
	handlers := contr.NewHandlers(db, *userContr, order)

	router := server.InitRoutes(handlers)
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}
	err := http.ListenAndServe(":"+PORT, router)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
