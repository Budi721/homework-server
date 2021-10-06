package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type Template struct {
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

func NewHttpResponse(r *http.Request, w http.ResponseWriter, httpCode int, result interface{}, err error) {
	if err != nil {
		Error(r, w, err, httpCode)
	} else {
		Success(w, result, httpCode)
	}
}

func Error(r *http.Request, w http.ResponseWriter, err error, httpCode int) {
	GenericError(r, w, err, err.Error(), httpCode)
}

func Success(w http.ResponseWriter, successResponse interface{}, responseCode ...int) {
	w.Header().Set("Content-Type", "application/json")

	if len(responseCode) == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(responseCode[0])
	}

	t := Template{
		Status: http.StatusOK,
		Result: successResponse,
		Error:  nil,
	}

	if successResponse != nil {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(t)
	}
}

func GenericError(r *http.Request, w http.ResponseWriter, err error, errorResponse interface{}, responseCode int) {
	if responseCode < 500 {
		log.Println(err.Error())
	} else {
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	t := Template{
		Status: responseCode,
		Result: nil,
		Error:  errorResponse,
	}

	if errorResponse != nil {
		_ = json.NewEncoder(w).Encode(t)
	}
}
