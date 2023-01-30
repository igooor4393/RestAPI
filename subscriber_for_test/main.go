package main

import (
	"RestAPI/pkg/logger"
	"fmt"
	"github.com/nats-io/nats.go"
)

var l = logger.Get()

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		l.Error().Err(err).Msg("Error connecting to nats server")
	}
	defer nc.Close()

	nc.Subscribe("requests", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Wait for messages
	select {}

}
