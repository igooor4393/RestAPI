package database

import (
	config "RestAPI/infrastructure"
	"RestAPI/logger"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var l = logger.Get()
var db *sql.DB

func Connect() {
	l.Info().Msg("Try login to the database")

	var err error
	db, err = sql.Open("postgres", config.LoadEnv())
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to the database")
		return
	}

	err = db.Ping()
	if err != nil {
		l.Error().Msg("Error: Could not establish a connection with the database")
		return
	}
}

func SaveRequest(requestType, input, output string) {
	// Save the request to the database
	l.Info().Msg("Save the database field to the database")

	_, err := db.Exec("INSERT INTO requests(requestType, input, output) VALUES ($1, $2, $3)", requestType, input, output)
	if err != nil {
		l.Error().Err(err).Msg("Error saving request to the database")
	}
}
