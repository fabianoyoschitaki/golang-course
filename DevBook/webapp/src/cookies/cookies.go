package cookies

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

// #IMPORTANT
// hashKey is required, used to authenticate values using HMAC. Create it using GenerateRandomKey().
// It is recommended to use a key with 32 or 64 bytes.
// blockKey is optional, used to encrypt values. Create it using GenerateRandomKey().
// The key length must correspond to the block size of the encryption algorithm.
// For AES, used by default, valid lengths are 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
// The default encoder used for cookie serialization is encoding/gob.
// Note that keys created using GenerateRandomKey() are not automatically persisted.
// New keys will be created when the application is restarted, and previously issued cookies will not be able to be decoded.
var s *securecookie.SecureCookie

// Configure creates SecureCookie using environment variables and should be called after having
func Configure(hashKey, blockKey []byte) {
	s = securecookie.New(hashKey, blockKey)
}

// ReadCookie reads the cookie DEVBOOK_DATA and returns a map with its contents
func ReadCookie(r *http.Request) (map[string]string, error) {
	cookie, error := r.Cookie("DEVBOOK_DATA")
	if error != nil {
		return nil, error
	}

	// here we have the cookie encoded and encrypted. Let's decode them to our map of values
	cookieValues := make(map[string]string)
	if error := s.Decode("data", cookie.Value, &cookieValues); error != nil {
		return nil, error
	}

	return cookieValues, nil
}

// DeleteCookie deletes cookie values in DEVBOOK_DATA
func DeleteCookie(rw http.ResponseWriter) {
	http.SetCookie(rw, &http.Cookie{
		Name:     "DEVBOOK_DATA",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0), // indicates it's expired
	})
}

// SaveCookie saves DEVBOOK_DATA cookie
func SaveCookie(rw http.ResponseWriter, ID, token string) error {
	// creating data to be saved into cookie
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	// let's encode and encrypt our data, this hapens due to hashKey and blockKey
	encodedData, error := s.Encode("data", data)
	if error != nil {
		return error
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "DEVBOOK_DATA",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}
