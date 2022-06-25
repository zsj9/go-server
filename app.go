package main

import (
	"net/http"
	"restful/routes"
	// "restful/controllers/user"
	// "restful/controllers/userinfo"
)

func main() {
	r := routes.NewRouter()

	// r.HandleFunc("/user/register", user.Register).Methods("POST")
	// r.HandleFunc("/user/login", user.Login).Methods("POST")

	// r.HandleFunc("/userinfo/add", userinfo.Add_userinfo).Methods("POST")
	// r.HandleFunc("/user/login", userinfo.Find_userinfo).Methods("POST")

	http.ListenAndServe(":8080", r)
}
