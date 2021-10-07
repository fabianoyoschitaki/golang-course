package authentication

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// go get github.com/dgrijalva/jwt-go
func CreateToken(ID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	// expiration for 6 hours
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // milliseconds after 01/01/1970, Unix epoch
	permissions["userId"] = ID

	// create token with signature method (not signed yet)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	// now let's sign the token with a secret
	return token.SignedString(config.SecretKey) // must be generated in a safe way. in the .env file
}
