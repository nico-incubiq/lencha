package controllers

import (
	"html/template"
	"net/http"

	"github.com/claisne/lencha/models"
	"github.com/claisne/lencha/session"
	"github.com/gorilla/mux"

	log "github.com/Sirupsen/logrus"
)

var problemsTemplates *template.Template

func CompileProblemTemplates() {
	problemsTemplates = template.Must(layout.Clone())
	problemsTemplates = template.Must(problemsTemplates.ParseGlob("./templates/problems/*.html"))
}

func Problem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problemName := vars["problem"]

	problem, err := models.GetProblemByName(problemName)
	if err != nil {
		http.Error(w, "Not found", http.StatusInternalServerError)
		return
	}

	params := struct {
		IsLogged        bool
		Problem         models.Problem
		DescriptionHTML template.HTML
	}{IsLogged: session.IsLogged(r), Problem: problem, DescriptionHTML: template.HTML(problem.Description)}

	err = problemsTemplates.ExecuteTemplate(w, "problem.html", params)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Template Error")
	}
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

	var problemsSolvedIds []int
	id, err := session.GetId(r)
	if err == nil {
		problemsSolvedIds, err = models.GetSolvedProblemsIdById(id)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Warn("Database error")

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if problemsSolvedIds == nil {
		problemsSolvedIds = []int{}
	}

	params := struct {
		IsLogged          bool
		Problems          []models.Problem
		ProblemsSolvedIds []int
	}{IsLogged: session.IsLogged(r), Problems: problems, ProblemsSolvedIds: problemsSolvedIds}

	err = problemsTemplates.ExecuteTemplate(w, "problems.html", params)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Template Error")
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
