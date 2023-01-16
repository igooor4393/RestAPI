package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func decrypt(w http.ResponseWriter, r *http.Request) {
	var req decryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement logic to decrypt the string here
	decrypted := req.Decrypt

	// Save the request to the database
	saveRequest("decrypt", req.Decrypt, decrypted)

	fmt.Fprintf(w, "Decrypted string: %s", decrypted)
}

func encrypt(w http.ResponseWriter, r *http.Request) {
	var req encryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement logic to encrypt the string here
	encrypted := req.Encrypt

	// Save the request to the database
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
	db, err := sql.Open("postgres", "postgres://user:password@host/database")
	if err != nil {

		log.Error().Msg("Error connecting to the database")
		//log.Print("Error connecting to the database: ", err)
		return

	}
	defer db.Close()

	// Save the request to the database
	_, err = db.Exec("INSERT INTO requests(type, input, output) VALUES ($1, $2, $3)", requestType, input, output)
	if err != nil {
		log.Error().Msg("Error inserting request into the database")
		return
	}
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Info().
			Str("From:", r.RemoteAddr).
			Str("Metod:", r.Method).
			Str("Request:", r.RequestURI)

		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/decrypt", middleware(decrypt))
	http.HandleFunc("/encrypt", middleware(encrypt))
	http.HandleFunc("/history", middleware(history))

	http.ListenAndServe(":8080", nil)

}
