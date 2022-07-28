package main

import (
	// "fmt"
	"net/http"
	"restful/routes"
)

func main () {
	router := routes.NewRouter();
	http.ListenAndServe(":8080", router)
}