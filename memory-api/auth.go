package main

import (
	"context"
	"net/http"
	"strings"
)

type AuthContext string

const (
	ContextCurrentUser AuthContext = "currentUser"
)

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")

		var currentUser = getClientProfileForToken(token)

		if currentUser == nil {
			http.Error(responseWriter, "forbidden", http.StatusForbidden)
			return
		}

		newContext := context.WithValue(request.Context(), ContextCurrentUser, currentUser)
		request = request.WithContext(newContext)

		next.ServeHTTP(responseWriter, request)
	}
}

func getClientProfileForToken(token string) *ClientProfile {
	if strings.HasPrefix(token, "Bearer ") {

		var extractedToken string = strings.TrimPrefix(token, "Bearer ")

		for _, profile := range database {
			if profile.Token == extractedToken {
				return &profile
			}
		}
	}

	return nil
}