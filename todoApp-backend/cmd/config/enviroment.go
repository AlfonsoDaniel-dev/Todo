package config

import (
	"github.com/joho/godotenv"
	"log"
)

func ReadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. ERROR %s", err)
	}

	verifyEnvVars()
}

func verifyEnvVars() {

	varsRequired := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME, APP_PORT. APP_HOST"}
	for _, v := range varsRequired {
		if v == "" {
			log.Fatalf("Required env variable %s is missing", v)
		}
	}
}
