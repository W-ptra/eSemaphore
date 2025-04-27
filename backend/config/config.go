package config

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	JWT_SECRET  []byte
	JWT_EXPIRED int
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loading env success")

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	jwtExpired,err := strconv.Atoi(os.Getenv("JWT_EXPIRED"))
	if err != nil {
		log.Fatal("Can't load or parse JWT_EXPIRED to int")
	}

	return Config{
		PORT: port,
		DB_HOST: dbHost,
		DB_PORT: dbPort,
		DB_USER: dbUser,
		DB_PASSWORD: dbPassword,
		DB_NAME: dbName,
		JWT_SECRET: jwtSecret,
		JWT_EXPIRED: jwtExpired,
	}
}