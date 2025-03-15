package main

import (
	"context"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var clientId = request.URL.Query().Get("clientId")
		clientProfile, ok := database[clientId]

		if !ok || clientId == "" {
			http.Error(responseWriter, "forbidden", http.StatusForbidden)
			return
		}

		token := request.Header.Get("Authorization")

		if !isValidToken(clientProfile, token) {
			http.Error(responseWriter, "forbidden", http.StatusForbidden)
			return
		}

		newContext := context.WithValue(request.Context(), "clientProfile", clientProfile)
		request = request.WithContext(newContext)
		next.ServeHTTP(responseWriter, request)
	}
}

func isValidToken(clientProfile ClientProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == clientProfile.Token
	}
	return false
}