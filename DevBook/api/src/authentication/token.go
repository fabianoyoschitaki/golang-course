package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// ValidateToken validates if request token is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	// let's convert it to JWT (header, claims and signature)
	token, error := jwt.Parse(tokenString, returnVerificationKey)
	if error != nil {
		return error
	}

	// token.Valid checks e.g. exp date. ok means we could get the claims
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token, sorry")
}

// returns if method being used to parse JWT is from the same family
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method! %v", token.Header["alg"])
	}
	// if it matches, let's return our secret key
	return config.SecretKey, nil
}

func extractToken(r *http.Request) string {
	// Bearer xyz
	token := r.Header.Get("Authorization")

	// validate format and returns the atual token
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// ExtractUserId returns user ID from token
func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, error := jwt.Parse(tokenString, returnVerificationKey)
	if error != nil {
		return 0, error
	}

	// extract userId from JWT claims
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		loggedUserID, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if error != nil {
			return 0, nil
		}
		return loggedUserID, nil
	}

	return 0, errors.New("invalid token, cannot extract userId claim")
}
