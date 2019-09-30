package controllers

import (
	"../serverGlobals"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Response interface {
	Status() int
	Error() error
	Body() interface{}
	Message() string
	Headers() map[string]string
}

type JsonResponse struct {
	Err_ error
	StatusCode_ int
	Message_ string
	Headers_ map[string]string
	Body_ interface{}
}

func (jr *JsonResponse) Error() error {
	return jr.Err_
}

func (jr *JsonResponse) Status() int {
	return jr.StatusCode_
}

func (jr *JsonResponse) Body() interface{} {
	return jr.Body_
}

func (jr *JsonResponse) Headers() map[string]string {
	return jr.Headers_
}

func (jr *JsonResponse) Message() string {
	return jr.Message_
}

type JsonResponseBaseJSON struct {
	Status int `json:"status"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type BaseController struct {
	SG *serverGlobals.ServerGlobals
}

func (bc *BaseController) WriteResponse(w http.ResponseWriter, response Response) {
	// Setting response Content-Type to `application/json` - default for JSON response
	w.Header().Set("Content-Type", "application/json")
	// Adding customer/additional Headers to response
	for key, value := range response.Headers() {
		w.Header().Set(key, value)
	}

	switch {
	case response.Status() >= 400 && response.Status() < 600:
		// Composing response message
		message := response.Message()
		if response.Error() != nil {
			message += response.Error().Error()
		}

		// Composing response container structure
		jsbJSON := JsonResponseBaseJSON{Status: response.Status(), Message: message, Data: nil}
		// JSON encoding
		jsonString, err := json.Marshal(jsbJSON)
		if err != nil {
			log.Println("can't encode response")
			w.WriteHeader(500)
			return
		}

		// Writing Headers back to client
		w.WriteHeader(response.Status())
		// Writing body back to client
		if _, err := io.WriteString(w, string(jsonString)); err != nil {
			panic("can't write response back to client")
		}
	case response.Status() >= 200 && response.Status() < 300:
		jsbJSON := JsonResponseBaseJSON{Status: response.Status(), Message: response.Message(), Data: response.Body()}
		jsonString, err := json.Marshal(jsbJSON)
		if err != nil {
			log.Println("can't encode response")
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(response.Status())
		if _, err := io.WriteString(w, string(jsonString)); err != nil {
			panic("can't write response back to client")
		}
	}
}