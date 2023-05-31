package utilities

import (
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {

	//  load the .env file
	godotenv.Load(".env")

	// get value from .env file
	value := os.Getenv(key)

	return value
}
