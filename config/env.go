package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host       string
	Port       string
	Address    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		Address: fmt.Sprintf("%s:%s", os.Getenv("Host"),
			os.Getenv("PORT")),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress: fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT")),
		DBName: os.Getenv("DB_NAME"),
	}
}
