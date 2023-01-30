package main

import (
	"RestAPI/handlers"
	"RestAPI/my_service/database"
	"RestAPI/my_service/middleware"
	"RestAPI/my_service/nats"
	"RestAPI/pkg/logger"
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
