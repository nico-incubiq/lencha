package controllers

import (
	"clem/lencha/models"
	"html/template"
	"net/http"

	log "github.com/Sirupsen/logrus"
	ctx "github.com/gorilla/context"
)

var usersTemplates *template.Template

func CompileUserTemplates() {
	usersTemplates = template.Must(layout.Clone()).Funcs(template.FuncMap{
		"sliceBy12": func(problems []models.Problem) [][]models.Problem {
			slices := make([][]models.Problem, 0, len(problems)/12+1)

			for i := 0; i+2 < len(problems); i += 12 {
				slice := make([]models.Problem, 12)
				slice[0] = problems[i]
				slice[1] = problems[i+1]
				slice[2] = problems[i+2]
				slices = append(slices, slice)
			}

			if len(problems)%12 != 0 {
				slice := make([]models.Problem, len(problems)%12)
				for i := 0; i < len(problems)%12; i++ {
					slice[i] = problems[len(problems)-len(problems)%12+i]
				}
				slices = append(slices, slice)
			}

			return slices
		},
	})
	usersTemplates = template.Must(usersTemplates.ParseGlob("./templates/users/*.html"))
}

func ApiUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Database error")
	}

	JSONResponse(w, users, http.StatusOK)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	user_ctx := ctx.Get(r, "user")
	if user_ctx == nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	user := user_ctx.(models.User)

	problemsSolved, err := user.GetSolvedProblems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := struct {
		IsLogged       bool
		User           models.User
		ProblemsSolved []models.Problem
	}{IsLogged: true, User: user, ProblemsSolved: problemsSolved}

	err = usersTemplates.ExecuteTemplate(w, "profile.html", params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
