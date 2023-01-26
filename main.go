package main

import (
	"RestAPI/handlers"
	"RestAPI/infrastructure"
	"RestAPI/infrastructure/database"
	"RestAPI/logger"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"net/http"
)

var l = logger.Get()

func init() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		l.Error().Err(err).Msg("Error connecting to nats server")
	}

	nc.Subscribe("requests", func(m *nats.Msg) {
		// Handle the received message here
		var response struct {
			RequestType string `json:"request_type,omitempty"`
			Input       string `json:"input"`
			Output      string `json:"output"`
		}
		json.Unmarshal(m.Data, &response)
		fmt.Printf("Received response: %+v\n", response)
	})
}

func main() {
	// load config files
	config.LoadEnv()

	// connect to database
	database.Connect()

	// connect to nats
	//nats.Open()

	// routing
	http.HandleFunc("/decrypt", handlers.Decrypt)
	http.HandleFunc("/encrypt", handlers.Encrypt)
	http.HandleFunc("/history", handlers.History)

	http.ListenAndServe(":8080", nil)
}
