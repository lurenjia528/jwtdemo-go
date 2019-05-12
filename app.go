package main

import (
	"github.com/lurenjia528/jwtdemo-go/routes"
	"net/http"
)

func main() {
	r := routes.NewRoute()

	http.ListenAndServe(":8080", r)
}
