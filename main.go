package main

import (
	"RestAPI/cryptLogic"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"os"

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

var Connection string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	Connection = fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func decrypt(w http.ResponseWriter, r *http.Request) {
	var req decryptRequest
	log.Info().Msg("get request for decrypt")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// logic to decrypt the string here
	decrypted := cryptLogic.Decod(req.Decrypt)

	// Save the request to the database
	log.Info().Msg("Save the decrypt request to the database")
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

	// logic to encrypt the string here
	encrypted := cryptLogic.Encode(req.Encrypt)

	// Save the request to the database
	log.Info().Msg("Save the encrypt request to the database")
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

	db, err := sql.Open("postgres", Connection)
	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, inspect all its fields, e.g.:
		log.Error().Msgf("pq error:%s", err.Code.Name())
		return
	}
	defer db.Close()

	// Save the request to the database
	log.Info().Msg("Save the data field to the database")

	_, err = db.Exec("INSERT INTO requests(requestType, input, output) VALUES ($1, $2, $3)", requestType, input, output)

	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, inspect all its fields, e.g.:
		log.Error().Msgf("pq error:%s", err.Code.Name())
		return
	}
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Info().
			Str("From:", r.RemoteAddr).
			Str("Metod:", r.Method).
			Str("Request:", r.RequestURI).
			Msg("Hi from Middleware")

		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/decrypt", middleware(decrypt))
	http.HandleFunc("/encrypt", middleware(encrypt))
	http.HandleFunc("/history", middleware(history))

	http.ListenAndServe(":8080", nil)

}
