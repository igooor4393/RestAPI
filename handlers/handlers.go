package handlers

import (
	"RestAPI/domain/cryptLogic"
	"RestAPI/infrastructure/database"

	"RestAPI/logger"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"

	"net/http"
)

var nc *nats.Conn
var l = logger.Get()

type decryptRequest struct {
	Decrypt string `json:"decrypt"`
}

type encryptRequest struct {
	Encrypt string `json:"encrypt"`
}

type historyRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func Decrypt(w http.ResponseWriter, r *http.Request) {
	var req decryptRequest
	l.Info().Msg("get request for decrypt")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		l.Error().Msg("Failed to decode the file")
		return
	}

	// logic to decrypt the string here
	decrypted := cryptLogic.Decod(req.Decrypt)

	// Save the request to the database

	// Публикатор отдает сообщение
	l.Info().Msg("Публикатор отдает сообщение")

	nc.Publish("requests", []byte(fmt.Sprintf("{requestType: %s, input: %s, output: %s}", "decrypt", req.Decrypt, decrypted)))

	l.Info().Msg("Save the decrypt request to the database")
	database.SaveRequest("decrypt", req.Decrypt, decrypted)

	fmt.Fprintf(w, "Decrypted string: %s", decrypted)
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	var req encryptRequest
	l.Info().Msg("get request for encrypt")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// logic to encrypt the string here
	encrypted := cryptLogic.Encode(req.Encrypt)

	// Save the request to the database
	l.Info().Msg("Save the encrypt request to the database")

	nc.Publish("requests", []byte(fmt.Sprintf("{requestType: %s, input: %s, output: %s}", "encypt", req.Encrypt, encrypted)))

	database.SaveRequest("encrypt", req.Encrypt, encrypted)

	fmt.Fprintf(w, "Encrypted string: %s", encrypted)
}

func History(w http.ResponseWriter, r *http.Request) {
	var req historyRequest
	l.Info().Msg("get request for history")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
