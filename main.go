package main

import (
	"RestAPI/handlers"
	"RestAPI/infrastructure"
	"RestAPI/infrastructure/database"
	"RestAPI/infrastructure/nats"

	"RestAPI/logger"
	"net/http"
)

var l = logger.Get()

func main() {
	// load config files
	config.LoadEnv()

	// connect to database
	database.Connect()

	// connect to nats
	nats.Open()

	// routing
	http.HandleFunc("/decrypt", handlers.Decrypt)
	http.HandleFunc("/encrypt", handlers.Encrypt)
	http.HandleFunc("/history", handlers.History)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		l.Error().Msgf("Error ListenAndServe: %s", err.Error())
	}
}
