package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	RedisHost  string
	RedisPort  string
	JWTSecret  string
	ServerPort string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.DBHost = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DBPort = cast.ToString(coalesce("DB_PORT", 5432))
	config.DBUser = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DBPassword = cast.ToString(coalesce("DB_PASSWORD", "password"))
	config.DBName = cast.ToString(coalesce("DB_NAME", "dbname"))

	config.RedisHost = cast.ToString(coalesce("REDIS_HOST", "localhost"))
	config.RedisPort = cast.ToString(coalesce("REDIS_PORT", ":6379"))

	config.JWTSecret = cast.ToString(coalesce("SMS_TOKEN", "my_secret"))

	config.ServerPort = cast.ToString(coalesce("SERVER_PORT", ":8080"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
