package service

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("asdasd")
	fmt.Fprintf(w, "this is homepage")
}
