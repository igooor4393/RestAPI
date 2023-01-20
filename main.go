package main

import (
	"RestAPI/cryptLogic"
	"RestAPI/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"time"

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

var l = logger.Get()

func init() {
	err := godotenv.Load()
	if err != nil {
		l.Error().Msg("Error loading .env file")
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
	l.Info().Msg("Save the decrypt request to the database")
	saveRequest("decrypt", req.Decrypt, decrypted)

	fmt.Fprintf(w, "Decrypted string: %s", decrypted)
}

func encrypt(w http.ResponseWriter, r *http.Request) {
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
	saveRequest("encrypt", req.Encrypt, encrypted)

	fmt.Fprintf(w, "Encrypted string: %s", encrypted)
}

func history(w http.ResponseWriter, r *http.Request) {
	var req historyRequest
	l.Info().Msg("get request for history")

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func saveRequest(requestType, input, output string) {

	// Connect to the database
	l.Info().Msg("Try login to the database")

	db, err := sql.Open("postgres", Connection)
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to the database")
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		l.Error().Msg("Error: Could not establish a connection with the database")
		return
	}

	// Save the request to the database
	l.Info().Msg("Save the data field to the database")

	_, err = db.Exec("INSERT INTO requests(requestType, input, output) VALUES ($1, $2, $3)", requestType, input, output)
	if err != nil {
		l.Error().Err(err).Msg("Error saving request to the database")
	}
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		l.
			Info().
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Str("user_agent", r.UserAgent()).
			Dur("elapsed_ms", time.Since(start)).
			Msg("incoming request")
		//log.Trace().
		//	Str("From:", r.RemoteAddr).
		//	Str("Metod:", r.Method).
		//	Str("Request:", r.RequestURI).
		//	Msg("Hi from Middleware!!!")

		next.ServeHTTP(w, r)
	}
}

func main() {

	addr := ":8080"
	l.Info().Msg("============================================================================================================================")
	l.Info().Msgf("Server started.")
	l.Info().Msgf("Server port - %s || BD port - %s", addr, os.Getenv("PORT"))

	http.HandleFunc("/decrypt", middleware(decrypt))
	http.HandleFunc("/encrypt", middleware(encrypt))
	http.HandleFunc("/history", middleware(history))

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		l.Error().Msgf("Error ListenAndServe: %s", err.Error())
	}

}
