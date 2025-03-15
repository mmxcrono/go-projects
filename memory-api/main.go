package main

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middleware{
	TokenAuthMiddleware,
}

func main() {
	log.Println("Starting application", "version", "1.0.0")
	setRoutes()
	startApiServer()
}

func setRoutes() {
	var handler http.HandlerFunc = handleClientProfile
	
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	http.HandleFunc("/user/profile", handler)
}

func startApiServer() {
	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println("Failed to start server")
		log.Println(err)
	}
}