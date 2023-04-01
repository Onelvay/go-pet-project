package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	rest "github.com/Onelvay/docker-compose-project/pkg/rest"
	"github.com/Onelvay/docker-compose-project/pkg/server"
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
	postgres := db.NewPostgresDb(*config)
	db := contr.NewDbController(postgres)

	handlers := rest.NewHandlers(db)

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
