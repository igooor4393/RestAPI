package nats

import (
	"RestAPI/logger"
	"encoding/json"
	"github.com/nats-io/nats.go"
)

var l = logger.Get()

var nc *nats.Conn

//func Open() (error, *nats.Conn) {
//	nc, err := nats.Connect("nats://localhost:4222")
//	if err != nil {
//		l.Error().Err(err).Msg("Error connecting to nats server")
//		return err, nil
//	}
//	return nil, nc
//}

func Publisher(requestType, input, output string) error {
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
	nc.Publish("requests", reqJson)
	return nil
}
