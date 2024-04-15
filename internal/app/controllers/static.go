package controllers

import (
	"github.com/JingBh/crypto-learn/internal/templates"
	"html/template"
	"net/http"
)

func ServeTemplate(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(templates.EmbedFS, path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
