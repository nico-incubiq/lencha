package main

import (
	"clem/lencha/config"
	"clem/lencha/controllers"
	"clem/lencha/models"
	"net/http"

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
