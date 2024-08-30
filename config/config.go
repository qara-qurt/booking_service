package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the application configuration
type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	Host     string // Host is the address of the database  - localhost
	Port     string // Port is the port of the database - 5432
	User     string // User is the username to connect to the database - postgres
	Password string // Password is the password to connect to the database - password
	SSLMode  string // SSLMode is the SSL mode to connect to the database - disable
	DBName   string // DBName is the name of the database - postgres
}

// GetEnv retrieves the value of an environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// New creates a new Config instance
func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	return &Config{
		Server: Server{
			Port: GetEnv("SERVER_PORT", "8080"),
		},
		Database: Database{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "5432"),
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "password"),
			DBName:   GetEnv("DB_NAME", "postgres"),
			SSLMode:  GetEnv("SSL_MODE", "disable"),
		},
	}, nil
}
