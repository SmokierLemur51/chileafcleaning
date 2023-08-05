package utils

import (
	"net/http"
	"html/template"

	"github.com/SmokierLemur51/chileafcleaning/models"
)

func RenderTemplate(w http.ResponseWriter, templateFile string, data models.PageData) error {
	tmpl, err := template.ParseFiles("templates/" + templateFile)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}