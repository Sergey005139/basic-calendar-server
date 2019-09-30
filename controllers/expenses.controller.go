package controllers

import (
	"../data"
	"../utils"
	"errors"
	"net/http"
)

type ExpensesController struct {
	BaseController
}

func (c ExpensesController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.RequestURI == "/expenses/list":
		c.WriteResponse(w,  c.list(w, r))

	case r.Method == http.MethodPost && r.RequestURI == "/expenses/add":
		c.WriteResponse(w, c.add(w, r))
	default:
		c.WriteResponse(w, &JsonResponse{StatusCode_: 404})
	}
}

func (c *ExpensesController) add(w http.ResponseWriter, r *http.Request) Response {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("authentication token empty")}
	}

	// Token decoding
	decodedToken, err := utils.Decode([]byte(authToken))
	if err != nil {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("can't decode authToken")}
	}

	u := c.SG.UserStorage.FindByUsername(string(decodedToken))
	if u == nil {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("invalid authentication token")}
	}

	expense := data.Expense{Date:"09-29-2019", Category:"Car", Amount:10.50}
	u.AddExpense(expense)

	return &JsonResponse{StatusCode_: http.StatusCreated, Message_: "expenses added", Body_: expense}
}

func (c *ExpensesController) list(w http.ResponseWriter, r *http.Request) Response  {
	// Getting Authorization token from Headers
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("authentication token empty")}
	}

	// Authentication token validation
	decodedToken, err := utils.Decode([]byte(authToken))
	if err != nil {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("can't decode authentication token")}
	}

	// Looking fro a user by Authentication Token
	u := c.SG.UserStorage.FindByUsername(string(decodedToken))
	if u == nil {
		return &JsonResponse{StatusCode_: http.StatusBadRequest, Err_: errors.New("invalid authentication token")}
	}

	return &JsonResponse{StatusCode_: http.StatusOK, Body_: u.Expenses}
}