package navigation

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type Film struct {
	Title string
	Director string
}

func (nr *navigationRouter) landingPage(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join(nr.templateBasePath, "web", "templates", "index.html")
	tmpl := template.Must(template.ParseFiles(templatePath))

	films := map[string][]Film{
		"Films": {
			{
				Title: "The Godfather",
				Director: "Francis Ford Coppola",
			},
			{
				Title: "Blad Runner",
				Director: "Ridley Scott",
			},
			{
				Title: "The Thing",
				Director: "John Carpenter",
			},
		},
	}

	tmpl.Execute(w, films)
}