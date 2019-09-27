package main

import (
	"./data"
	"./controllers"
	"./serverGlobals"
	"log"
	"net/http"
)



func main() {
	sg := serverGlobals.ServerGlobals{UserStorage: data.UserStorage{}}

	http.Handle("/auth/", controllers.AuthController{SG: &sg})
	log.Fatal(http.ListenAndServe(":8080", nil))
}