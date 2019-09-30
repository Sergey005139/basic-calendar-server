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
	baseController := controllers.BaseController{SG: &sg}

	http.Handle("/auth/", controllers.AuthController{BaseController: baseController})
	http.Handle("/expenses/", controllers.ExpensesController{BaseController: baseController})
	log.Fatal(http.ListenAndServe(":8080", nil))
}