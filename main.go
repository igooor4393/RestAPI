package main

import (
	"RestAPI/cryptLogic"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"net/http"

	_ "github.com/lib/pq"
)

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

const (
	host     = "localhost"
	port     = 32768
	user     = "postgres"
	password = "postgrespw"
	dbname   = "postgres"
)

func decrypt(w http.ResponseWriter, r *http.Request) {
	var req decryptRequest
	log.Info().Msg("get request for decrypt")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement logic to decrypt the string here
	decrypted := cryptLogic.Decod(req.Decrypt)

	// Save the request to the database
	log.Info().Msg("Try save the decrypt request to the database")
	saveRequest("decrypt", req.Decrypt, decrypted)

	fmt.Fprintf(w, "Decrypted string: %s", decrypted)
}

func encrypt(w http.ResponseWriter, r *http.Request) {
	var req encryptRequest
	log.Info().Msg("get request for encrypt")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement logic to encrypt the string here
	encrypted := cryptLogic.Encode(req.Encrypt)

	// Save the request to the database
	log.Info().Msg("Try save the encrypt request to the database")
	saveRequest("encrypt", req.Encrypt, encrypted)

	fmt.Fprintf(w, "Encrypted string: %s", encrypted)
}

func history(w http.ResponseWriter, r *http.Request) {
	var req historyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func saveRequest(requestType, input, output string) {
	// Connect to the database
	log.Info().Msg("Try login to the database")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, inspect all its fields, e.g.:
		log.Error().Msgf("pq error:", err.Code.Name())
		return
	}
	defer db.Close()

	// Save the request to the database
	log.Info().Msg("Try Save the data field to the database")

	_, err = db.Exec("INSERT INTO requests(requestType, input, output) VALUES ($1, $2, $3)", requestType, input, output)

	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, inspect all its fields, e.g.:
		log.Error().Msgf("pq error:", err.Code.Name())
		return
	}
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Info().
			Str("From:", r.RemoteAddr).
			Str("Metod:", r.Method).
			Str("Request:", r.RequestURI).
			Msg("Hi from middleware")

		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/decrypt", middleware(decrypt))
	http.HandleFunc("/encrypt", middleware(encrypt))
	http.HandleFunc("/history", middleware(history))

	http.ListenAndServe(":8080", nil)

}
