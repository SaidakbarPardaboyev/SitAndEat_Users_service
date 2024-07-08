package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT         string
	DB_HOST           string
	DB_PORT           string
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	SIGNING_KEY       string
  }
  
func Load() Config{
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	
	config := Config{}
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "postgres"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "postgres"))
	config.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", "50051"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "secret"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}