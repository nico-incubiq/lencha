package main

import (
	"clem/lencha/config"
	"clem/lencha/controllers"
	"clem/lencha/middlewares"
	"clem/lencha/problems"
	"net/http"

	"github.com/gorilla/mux"
)

func GenerateHandlers() http.Handler {
	var handlers http.Handler

	routesHandlers := CreateRoutes()

	if !config.Conf.Production {
		handlers = middlewares.LoggingRequest(routesHandlers)
	}

	return handlers
}

func CreateRoutes() http.Handler {
	// Router definition
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Home
	router.Handle("/", http.HandlerFunc(controllers.Home))

	// Problems
	problemsRouter := router.Path("/problems").Subrouter()
	problemsRouter.Methods("GET").HandlerFunc(controllers.Problems)

	// Problems Api
	problemsApiRouter := apiRouter.PathPrefix("/problems").Subrouter()
	problemsApiRouter.HandleFunc("/", controllers.ApiProblems)
	handlerReverse := middlewares.RequireApiKey(problems.HandlerFromStateHandler(problems.Reverse))
	problemsApiRouter.HandleFunc("/reverse", handlerReverse)

	// Profile
	handlerProfile := middlewares.RequireLogged(http.HandlerFunc(controllers.Profile))
	router.Path("/profile").Methods("GET").HandlerFunc(handlerProfile)

	// User Api
	usersApiRouter := apiRouter.Path("/users").Subrouter()
	usersApiRouter.Methods("GET").HandlerFunc(controllers.ApiUsers)

	// Login
	router.Path("/login").Methods("POST").HandlerFunc(controllers.LoginPost)

	// Logout
	router.Path("/logout").Methods("GET").HandlerFunc(controllers.Logout)

	// Register
	router.Path("/register").Methods("POST").HandlerFunc(controllers.RegisterPost)

	// Static assets
	router.PathPrefix("/fonts").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./public/fonts/"))))
	router.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css/"))))
	router.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js/"))))

	// 404
	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/404.html")
}
