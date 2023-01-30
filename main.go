package main

import (
	"RestAPI/handlers"
	"RestAPI/infrastructure/database"
	"RestAPI/infrastructure/nats"
	"RestAPI/logger"
	"RestAPI/my_service/middleware"
	"net/http"
)

var l = logger.Get()

func main() {
	// load config files

	// connect to database
	database.Connect()

	// connect to nats
	nats.Open()

	// routing
	http.HandleFunc("/decrypt", middleware.Midd(handlers.Decrypt))
	http.HandleFunc("/encrypt", middleware.Midd(handlers.Encrypt))
	http.HandleFunc("/history", middleware.Midd(handlers.History))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		l.Error().Msgf("Error ListenAndServe: %s", err.Error())
	}
}
