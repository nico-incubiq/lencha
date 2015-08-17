package main

import (
	"net/http"

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

	http.ListenAndServe(config.Conf.Host, GenerateHandlers())
	log.WithFields(log.Fields{"host": config.Conf.Host}).Info("Server started")
}
