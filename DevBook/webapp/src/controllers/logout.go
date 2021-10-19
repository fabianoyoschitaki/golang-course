package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// DoLogout does user logout by removing cookies
func DoLogout(rw http.ResponseWriter, r *http.Request) {
	cookies.DeleteCookie(rw)
	http.Redirect(rw, r, "/login", http.StatusFound)
}
