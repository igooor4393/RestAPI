package main

import (
	"RestAPI/handlers"
	"RestAPI/my_service/database"
	"RestAPI/my_service/middleware"
	"RestAPI/my_service/nats"
	"RestAPI/my_service/ticker"
	"RestAPI/pkg/logger"
	"net/http"
)

var l = logger.Get()

func main() {

	// connect to database
	database.Connect()

	// connect to nats
	nats.Open()

	//Раз в 30 секунд проверяет соединение с NATS и BD. Конечный таймер поставлен на 30 минут, через 30 минут тикер перестает срабатывать.
	go ticker.Ticker()

	// routing
	http.HandleFunc("/decrypt", middleware.Midd(handlers.Decrypt))
	http.HandleFunc("/encrypt", middleware.Midd(handlers.Encrypt))
	http.HandleFunc("/history", middleware.Midd(handlers.History))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		l.Error().Msgf("Error ListenAndServe: %s", err.Error())
	}
}
