package middlewares

import "net/http"

// middlewares contains layers between request and response. Often used to something that should be appied to ALL routes
func Authenticate(next http.HandlerFunc) http.HandlerFunc {

}
