package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is the base url for our backend API
	APIURL = ""

	// Port is where web application is running
	ApplicationPort = 0

	// HashKey is used to authenticate the cookie
	HashKey []byte

	// BlockKey is used to encrypt cookie data
	BlockKey []byte
)

// LoadConfiguration initializes environment variables
func LoadConfiguration() {
	var error error

	// Tries to load environment variables
	if error = godotenv.Load(); error != nil {
		log.Fatal(error) // if we have issues related to our environment variables, we cannot continue
	}

	// ApplicationPort
	ApplicationPort, error = strconv.Atoi(os.Getenv("APP_PORT"))
	if error != nil {
		ApplicationPort = 9000
	}

	// APIURL
	APIURL = os.Getenv("API_URL")

	// HashKey
	HashKey = []byte(os.Getenv("HASH_KEY"))

	// HashKey
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
