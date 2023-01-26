package config

import (
	"RestAPI/logger"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var l = logger.Get()

//var connection string

func LoadEnv() string {
	err := godotenv.Load()
	if err != nil {
		l.Error().Msg("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	connection := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return connection
}
