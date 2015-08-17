package controllers

import (
	"html/template"
	"net/http"

	"github.com/claisne/lencha/middlewares"
)

var homeTemplates *template.Template

func CompileHomeTemplates() {
	homeTemplates = template.Must(homeTemplates.ParseGlob("./templates/*.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {

	params := struct {
		IsLogged bool
	}{IsLogged: middlewares.IsLogged(r)}

	err := homeTemplates.ExecuteTemplate(w, "index.html", params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
