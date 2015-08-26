package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/claisne/lencha/config"
	"github.com/claisne/lencha/controllers"
	"github.com/claisne/lencha/models"

	log "github.com/Sirupsen/logrus"
)

func main() {
	// Setup global vars
	config.Load()
	models.ConnectDb()
	models.ConnectRedis()
	controllers.CompileTemplates()

	// Init random seed for the problems
	rand.Seed(time.Now().UTC().UnixNano())

	log.WithFields(log.Fields{"host": config.Conf.Host}).Info("Starting server")
	err := http.ListenAndServe(config.Conf.Host, GenerateHandlers())
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Warn("Error starting server")
	}
}
