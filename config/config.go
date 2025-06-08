package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	PasetoSymmKey string
	ServerPort    string
	DatabaseURL   string
}

func Load(path string) (*Config, error) {
	godotenv.Load(path + "/.env")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	pasetoKey := os.Getenv("PASETO_SYMMETRIC_KEY")
	serverPort := os.Getenv("SERVER_PORT")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	return &Config{
		DBHost:        dbHost,
		DBPort:        dbPort,
		DBUser:        dbUser,
		DBPassword:    dbPass,
		DBName:        dbName,
		PasetoSymmKey: pasetoKey,
		ServerPort:    serverPort,
		DatabaseURL:   dbURL,
	}, nil
}
