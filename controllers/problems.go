package controllers

import (
	"clem/lencha/middlewares"
	"clem/lencha/models"
	"html/template"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

var problemsTemplates *template.Template

func CompileProblemTemplates() {
	problemsTemplates = template.Must(layout.Clone())
	problemsTemplates = template.Must(problemsTemplates.ParseGlob("./templates/problems/*.html"))
}

func Problems(w http.ResponseWriter, r *http.Request) {
	problems, err := models.GetAllProblems()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Database error")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := struct {
		IsLogged bool
		Problems []models.Problem
	}{IsLogged: middlewares.IsLogged(r), Problems: problems}

	err = problemsTemplates.ExecuteTemplate(w, "index.html", params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ApiProblems(w http.ResponseWriter, r *http.Request) {
	problems, err := models.GetAllProblems()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Database error")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, problems, http.StatusOK)
}
