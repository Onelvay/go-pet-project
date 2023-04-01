package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "io/ioutil"

	mdl "github.com/Onelvay/docker-compose-project/pkg/model"
	"github.com/Onelvay/docker-compose-project/pkg/server"
	"github.com/spf13/viper"

	// yaml "gopkg.in/yaml.v3"
	db "github.com/Onelvay/docker-compose-project/database"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	config := mdl.NewConfig(viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.dbname"),
		viper.GetString("db.user"),
		viper.GetString("db.pass"),
	)
	db.NewPostgresDb(*config)

	router := server.InitRoutes()
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
