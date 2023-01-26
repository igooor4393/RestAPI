package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to messages on the "orders" topic
	_, err = nc.Subscribe("requests", func(msg *nats.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	})
	if err != nil {
		log.Println(err)
	}

	// Wait for messages
	select {}
}
