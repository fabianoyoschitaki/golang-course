package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// generic interface. to be used by any kind of data, we need interface{}
// JSON writes json to response
func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	// only if data is valid we should write to response
	if data != nil {
		// writes data JSON to response
		if error := json.NewEncoder(rw).Encode(data); error != nil {
			log.Fatal(error)
		}
	}
}

// Error writes error to response
func Error(rw http.ResponseWriter, statusCode int, e error) {
	JSON(rw, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: e.Error(), // error message
	})
}
