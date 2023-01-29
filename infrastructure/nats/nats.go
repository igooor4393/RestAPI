package nats

import (
	"RestAPI/logger"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
)

var l = logger.Get()

func Open() (error, *nats.Conn) {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		l.Error().Err(err).Msg("Error connecting to nats server")
		return err, nil
	}

	return nil, nc
}

func Publisher(nc *nats.Conn, requestType, input, output string) error {

	defer nc.Close()

	req := struct {
		RequestType string `json:"requestType"`
		Input       string `json:"input"`
		Output      string `json:"output"`
	}{requestType, input, output}

	reqJson, err := json.Marshal(req)

	if err != nil {
		l.Error().Err(err).Msg("Error marshalling request")
		return err
	}
	fmt.Println("37")
	nc.Publish("requests", reqJson)
	return nil
}

//func Publisher(requestType, input, output string) error {
//	nc, err := nats.Connect("nats://localhost:4222")
//	if err != nil {
//		l.Error().Err(err).Msg("Error connecting to nats server")
//	}
//	defer nc.Close()
//
//	req := struct {
//		RequestType string `json:"requestType"`
//		Input       string `json:"input"`
//		Output      string `json:"output"`
//	}{requestType, input, output}
//
//	reqJson, err := json.Marshal(req)
//
//	if err != nil {
//		l.Error().Err(err).Msg("Error marshalling request")
//		return err
//	}
//	fmt.Println("37")
//	nc.Publish("requests", reqJson)
//	return nil
//}
