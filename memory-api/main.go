package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.Default()
	logger.Info("Starting application", "version", "1.0.0")
	setRoutes()
	startApiServer()
}

func setRoutes() {
	http.HandleFunc("/user/profile", handleClientProfile)
}

func startApiServer() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		logger.Error("Failed to start server", "err", err)
	}
}