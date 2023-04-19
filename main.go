package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Onelvay/docker-compose-project/payment/client"
	contr "github.com/Onelvay/docker-compose-project/pkg/controller"
	"github.com/Onelvay/docker-compose-project/pkg/server"
	"github.com/Onelvay/docker-compose-project/pkg/service"
	db "github.com/Onelvay/docker-compose-project/postgres"
	"github.com/spf13/viper"
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
	hasher := service.NewHasher(viper.GetString("app.hash"))
	order := contr.NewOrderDbController(postgres)

	userContr := contr.NewUserController(userDb, tokenDb, hasher, order)
	handlers := contr.NewHandlers(db, &userContr, order, tokenDb)

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
