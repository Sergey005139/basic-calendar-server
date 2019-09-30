package controllers

import (
	"../utils"
	"../data"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type AuthController struct {
	BaseController
	counter int
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
	c.RW = w
	c.R = r

	switch {
	case r.Method == http.MethodPost && r.RequestURI == "/auth/signup":
		c.WriteResponse(c.signup())
	case r.Method == http.MethodPost && r.RequestURI == "/auth/signin":
		c.WriteResponse(c.signin())
	default:
		c.WriteResponse(&JsonResponse{StatusCode_: 404, Message_: "not found"})
	}
}

// Responses:
// - 400 - Login or password incorrect
func (c *AuthController) signin() Response {
	var ds SignupDS
	if err := json.NewDecoder(c.R.Body).Decode(&ds); err != nil {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("can't json decode incoming payload")}
	}

	if user := c.SG.UserStorage.FindByUsername(ds.Username); user != nil {
		return &JsonResponse{
			StatusCode_: http.StatusOK,
			Err_: errors.New("can't json decode incoming payload"),
			Headers_: map[string]string{"Authentication": string(utils.Encode([]byte(user.Username)))},
			Body_: user,
		}
	} else {
		log.Println("not found!")
		return &JsonResponse{StatusCode_: http.StatusNotFound, Message_: "invalid credentials"}
	}
}

func (c *AuthController) signup() Response {
	var ds SignupDS
	if err := json.NewDecoder(c.R.Body).Decode(&ds); err != nil {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("can't json decode incoming request")}
	}

	u := data.User{Username: ds.Username, Password: ds.Password}

	if c.SG.UserStorage.IsExist(&u) {
		return &JsonResponse{StatusCode_: http.StatusForbidden, Err_: errors.New("user already exist")}
	}

	if ok, _ := c.SG.UserStorage.Insert(&u); !ok {
		return &JsonResponse{StatusCode_: http.StatusInternalServerError, Err_: errors.New("can't create user")}
	}

	return &JsonResponse{StatusCode_: 201, Message_: "user created"}
}