package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func GetEnv(key, def string) string {
	// First check environment variables
	if val := os.Getenv(key); val != "" {
		return val
	}
	// Then check .env file
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	envFile := ".env"
	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		// Don't panic if .env file doesn't exist, just use empty map
		// This allows using environment variables without .env file
		Env = make(map[string]string)
	}
}
