package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct {
	Code int
	Message string
}

func writeError(responseWriter http.ResponseWriter, message string, code int) {
	response := Error {
		Code: code,
		Message: message,
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(code)

	json.NewEncoder(responseWriter).Encode(response)
}

var (
	RequestErrorHandler = func(responseWriter http.ResponseWriter, err error) {
		writeError(responseWriter, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHanddler = func(responseWriter http.ResponseWriter) {
		writeError(responseWriter, "An unexpected error occurred", http.StatusInternalServerError)
	}
)