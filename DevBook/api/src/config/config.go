package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseConnectionString is the connection string for MySQL
	DatabaseConnectionString = ""

	// Port where the API is running
	ApplicationPort = 0

	// Key used to sign JWT
	SecretKey []byte
)

// LoadConfiguration initializes environment variables
func LoadConfiguration() {
	var error error

	// Tries to load environment variables
	if error = godotenv.Load(); error != nil {
		log.Fatal(error) // if we have issues related to our environment variables, we cannot continue
	}

	// Application port
	ApplicationPort, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		ApplicationPort = 9000
	}

	// MySQL connection string
	DatabaseConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// JWT secret key
	SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}
