package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

var layout *template.Template

func CompileTemplates() {
	layout = template.Must(template.ParseFiles("./templates/layout.html"))
	CompileHomeTemplates()
	CompileUserTemplates()
	CompileProblemTemplates()
}

func JSONResponse(w http.ResponseWriter, v interface{}, c int) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Info("Error creating JSON response")
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	w.Write(jsonBytes)
}
