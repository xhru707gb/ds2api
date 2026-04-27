package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Retrieve configuration from environment
	port := getEnv("PORT", "3000") // changed default from 8080 to 3000 to avoid conflicts on my machine
	dsHost := getEnv("DS_HOST", "")
	dsPort := getEnv("DS_PORT", "5001") // my NAS uses 5001 (HTTPS) instead of 5000 (HTTP)
	dsUser := getEnv("DS_USER", "")
	dsPass := getEnv("DS_PASS", "")

	if dsHost == "" {
		log.Fatal("DS_HOST environment variable is required")
	}
	if dsUser == "" {
		log.Fatal("DS_USER environment variable is required")
	}
	if dsPass == "" {
		log.Fatal("DS_PASS environment variable is required")
	}

	cfg := &Config{
		Port:   port,
		DSHost: dsHost,
		DSPort: dsPort,
		DSUser: dsUser,
		DSPass: dsPass,
	}

	server := NewServer(cfg)

	log.Printf("Starting ds2api server on port %s", port)
	log.Printf("Connecting to Synology at %s:%s", dsHost, dsPort)
	log.Printf("Build: personal fork - https://github.com/me/ds2api")
	if err := server.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// getEnv retrieves an environment variable or returns a fallback default value.
// If the key exists but is set to an empty string, the fallback is returned instead.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
