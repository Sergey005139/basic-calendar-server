package controllers

import (
	"../data"
	"../serverGlobals"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthController struct {
	SG *serverGlobals.ServerGlobals
}

type SignupDS struct {
	 Username string
	 Password string
}

type SigninDS struct {
	Username string
	Password string
}

func (c AuthController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && r.RequestURI == "/auth/signup":
		c.signup(w, r)
	case r.Method == http.MethodPost && r.RequestURI == "/auth/signin":
		c.signin(w, r)
	default:
		w.WriteHeader(404)
		fmt.Fprint(w,"Not Found")

	}
}

//
// Responses:
// - 400 - Login or password incorrect
func (c *AuthController) signin(w http.ResponseWriter, r *http.Request) {
	var ds SignupDS
	if err := json.NewDecoder(r.Body).Decode(&ds); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w,err)
	}

	if user := c.SG.UserStorage.FindByUsername(ds.Username); user != nil {
		fmt.Fprint(w, "Exist")
	} else {
		fmt.Fprint(w, "Not Exist")
	}
}

func (c *AuthController) signup(w http.ResponseWriter, r *http.Request) {
	var ds SignupDS
	if err := json.NewDecoder(r.Body).Decode(&ds); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w,err)
	}

	u := data.User{Username: ds.Username, Password: ds.Password}

	c.SG.UserStorage.Insert(&u)

	fmt.Printf("%#v", c.SG.UserStorage)

	fmt.Fprintln(w, "SignUp")
}