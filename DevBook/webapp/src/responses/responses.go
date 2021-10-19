package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// APIError represents error response
type APIError struct {
	Error string `json:"error"`
}

func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	// #IMPORTANT http: request method or response status code does not allow body
	// This is because data might be nil
	// if data != nil { this is a bad check

	if statusCode != http.StatusNoContent {
		// writes data JSON to response
		if error := json.NewEncoder(rw).Encode(data); error != nil {
			log.Fatal(error)
		}
	}
}

// HandleHttpResponseErrors handles requests with status codes >= 400 (errors)
func HandleHttpResponseErrors(rw http.ResponseWriter, r *http.Response) {
	var error APIError
	json.NewDecoder(r.Body).Decode(&error)
	// now we forward the backend API to FE API
	JSON(rw, r.StatusCode, error)
}
