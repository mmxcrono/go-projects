package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mmxcrono/go-projects/memory-api/internal/auth"
	"github.com/mmxcrono/go-projects/memory-api/internal/db"
)

func HandleClientProfile(responseWriter http.ResponseWriter, request *http.Request) {
	switch (request.Method) {
	case http.MethodGet:
		GetClientProfile(responseWriter, request)
	case http.MethodPatch:
		UpdateClientProfile(responseWriter, request)
	default:
		http.Error(responseWriter, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(responseWriter http.ResponseWriter, request *http.Request) {
	var clientId = strings.TrimSpace(request.URL.Query().Get("clientId"))
	
	if clientId == "" {
		log.Println("clientId is required")
		http.Error(responseWriter, "clientId is required", http.StatusBadRequest)
		return
	}

	clientProfile, ok := db.Database[clientId]

	if !ok {
		log.Printf("no client found for clientId %v\n", clientId)
		http.Error(responseWriter, "not found", http.StatusNotFound)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")

	response := db.ClientProfile{
		Email: clientProfile.Email,
		Name: clientProfile.Name,
		Id: clientProfile.Id,
	}

	json.NewEncoder(responseWriter).Encode(response)
}

func UpdateClientProfile(responseWriter http.ResponseWriter, request *http.Request) {
	currentUser := request.Context().Value(auth.ContextCurrentUser).(*db.ClientProfile)

	var clientId = strings.TrimSpace(request.URL.Query().Get("clientId"))
	
	if clientId == "" || currentUser.Id != clientId {
		log.Printf("clientId does not match current user %v\n", clientId)
		http.Error(responseWriter, "forbidden", http.StatusForbidden)
		return
	}

	clientProfile, ok := db.Database[clientId]

	if !ok {
		log.Printf("no client found for clientId %v\n", clientId)
		http.Error(responseWriter, "not found", http.StatusNotFound)
		return
	}

	var payloadData db.ClientProfile

	if err := json.NewDecoder(request.Body).Decode(&payloadData); err != nil {
		log.Printf("invalid JSON\n")
		http.Error(responseWriter, "invalid JSON", http.StatusBadRequest)
		return
	}

	defer request.Body.Close()

	if payloadData.Email != "" {
		clientProfile.Email = payloadData.Email
	}
	
	if payloadData.Name != "" {
		clientProfile.Name = payloadData.Name
	}

	db.Database[clientProfile.Id] = clientProfile

	responseWriter.WriteHeader(http.StatusOK)
}
