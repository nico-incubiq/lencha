package controllers

import (
	"html/template"
	"net/http"

	"github.com/claisne/lencha/session"

	log "github.com/Sirupsen/logrus"
)

var homeTemplates *template.Template

func CompileHomeTemplates() {
	homeTemplates = template.Must(homeTemplates.ParseGlob("./templates/*.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {

	params := struct {
		IsLogged bool
	}{IsLogged: session.IsLogged(r)}

	err := homeTemplates.ExecuteTemplate(w, "index.html", params)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Template Error")
	}
}
