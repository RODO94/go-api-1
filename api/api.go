package api

import (
	"encoding/json"
	"net/http"
)

// This type defines the params in a Coin Balance Request
type CoinBalanceParams struct {
	Username string
}


// This type defines the Response for a Coin Balance Request
type CoinBalanceResponse struct {
	// HTTP status code
	Code int 
	Balance int64
}

type Error struct{
	Code int
	Message string
}

// This function creates and sends an error
func writeError(w http.ResponseWriter, message string, code int){

resp := Error{
	Code: code,
	Message: message,
}

w.Header().Set("Content-Type","application/json")

w.WriteHeader(code)

json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Error Occured",http.StatusInternalServerError)
	}
)