package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	RedisHost     string
	RedisPort     string
	JWTSecret     string
	ServerPort    string
	GroupId       string
	BotToken      string
	EmailName     string
	EmailPassword string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		GroupId:       os.Getenv("GROUP_ID"),
		BotToken:      os.Getenv("BOT_TOKEN"),
		EmailName:     os.Getenv("EMAIL_NAME"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
	}
}
