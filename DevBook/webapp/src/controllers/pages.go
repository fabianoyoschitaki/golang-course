package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginPage renders login page LOL
func LoadLoginPage(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "login.html", nil)
}

// LoadSignUpPage renders create user page
func LoadSignUpPage(rw http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(rw, "signup.html", nil)
}
