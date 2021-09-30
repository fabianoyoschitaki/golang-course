package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// write outside of main function as a convention
var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {
	// all templates we're going to create (*.html files in the same folder)
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// creating user to be displayed at home.html
		u := user{"John", "john@doe.com"}

		// more than 1 template...
		templates.ExecuteTemplate(w, "home.html", u) // the writer, the template and properties to be replaced in the final html
	})
	fmt.Println("Listening at port :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
