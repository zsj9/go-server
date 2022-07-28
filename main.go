package main

import (
	// "fmt"
	"net/http"
	"restful/routes"
)

func main () {
	router := routes.NewRouter();
	http.ListenAndServe(":3000", router)
}