package handlers

import (
	"net/http"
	"html/template"
	"fmt"

	"github.com/SmokierLemur51/goleafdev/data"
	"github.com/SmokierLemur51/goleafdev/api"

	"github.com/go-chi/chi/v5"
)

var ENDPOINTS data.PublicEndPoints = data.PublicEndPoints{Index: "/", About: "/about", Services: "/services"}


func RenderTemplate(w http.ResponseWriter, data data.PublicPageData) error {
	tmpl, err := template.ParseFiles("templates/" + data.Page)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}


func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := api.CallEmailScript()
	CheckErr(err)
	
	page := data.PublicPageData{
		Page: "index.html",
		Title: "Greenleaf Development",
		Content: "Sample Content",
		Endpoints: ENDPOINTS,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func ServicesHandler(w http.ResponseWriter, r *http.Request) {

	page := data.PublicPageData{
		Page: "services.html",
		Title: "Greenleaf Development",
		Content: "Sample Content",
		Endpoints: ENDPOINTS,
	}

	err := RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AdboutHandler(w http.ResponseWriter, r *http.Request) {

	page := data.PublicPageData{
		Page: "about.html",
		Title: "Greenleaf Development",
		Content: "Sample Content",
		Endpoints: ENDPOINTS,
	}

	err := RenderTemplate(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func SendEmails(w http.ResponseWriter, r *http.Request) {
	target := chi.URLParam(r, "target")
	subject := chi.URLParam(r, "subject")
	message := chi.URLParam(r, "message")

	fmt.Printf("Sending Message to %s:\n\tSubject: %s\n\t\t%s", target, subject, message)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}