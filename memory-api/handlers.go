package main

import (
	"encoding/json"
	"net/http"
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
	// Get from context, set by middleware
	clientProfile := request.Context().Value("clientProfile").(ClientProfile)

	responseWriter.Header().Set("Content-Type", "application/json")

	response := ClientProfile{
		Email: clientProfile.Email,
		Name: clientProfile.Name,
		Id: clientProfile.Id,
	}

	json.NewEncoder(responseWriter).Encode(response)
}

func UpdateClientProfile(responseWriter http.ResponseWriter, request *http.Request) {
	// Get from context, set by middleware
	clientProfile := request.Context().Value("clientProfile").(ClientProfile)

	var payloadData ClientProfile

	if err := json.NewDecoder(request.Body).Decode(&payloadData); err != nil {
		http.Error(responseWriter, "invalid JSON", http.StatusBadRequest)
		return
	}

	defer request.Body.Close()

	clientProfile.Email = payloadData.Email
	clientProfile.Name = payloadData.Name
	database[clientProfile.Id] = clientProfile

	responseWriter.WriteHeader(http.StatusOK)
}
