package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func handleClientProfile(responseWriter http.ResponseWriter, request *http.Request) {
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

	clientProfile, ok := database[clientId]

	if !ok {
		log.Printf("no client found for clientId %v\n", clientId)
		http.Error(responseWriter, "not found", http.StatusNotFound)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")

	response := ClientProfile{
		Email: clientProfile.Email,
		Name: clientProfile.Name,
		Id: clientProfile.Id,
	}

	json.NewEncoder(responseWriter).Encode(response)
}

func UpdateClientProfile(responseWriter http.ResponseWriter, request *http.Request) {
	currentUser := request.Context().Value(ContextCurrentUser).(*ClientProfile)

	var clientId = strings.TrimSpace(request.URL.Query().Get("clientId"))
	
	if clientId == "" || currentUser.Id != clientId {
		log.Printf("clientId does not match current user %v\n", clientId)
		http.Error(responseWriter, "forbidden", http.StatusForbidden)
		return
	}

	clientProfile, ok := database[clientId]

	if !ok {
		log.Printf("no client found for clientId %v\n", clientId)
		http.Error(responseWriter, "not found", http.StatusNotFound)
		return
	}

	var payloadData ClientProfile

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

	database[clientProfile.Id] = clientProfile

	responseWriter.WriteHeader(http.StatusOK)
}
