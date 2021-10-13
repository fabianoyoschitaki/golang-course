package utils

import (
	"net/http"
	"text/template"
)

// This file interacts with go html/templates.
// We have 2 responsibilites:
// 1. to load all html into templates variable
// 2. to render the pages

var templates *template.Template

// LoadTemplates inserts html templates inside templates variable
func LoadTemplates() {
	// pointing where our HTML files are
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// RenderTemplate renders a specific template with actual data
func RenderTemplate(rw http.ResponseWriter, templateName string, data interface{}) {
	templates.ExecuteTemplate(rw, templateName, data)
}
