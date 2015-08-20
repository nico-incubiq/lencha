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
	http.ListenAndServe(config.Conf.Host, GenerateHandlers())
}
