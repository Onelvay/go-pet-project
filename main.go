package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Onelvay/docker-compose-project/pkg/server"
)

func main() {
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
