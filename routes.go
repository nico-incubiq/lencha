package main

import (
	"net/http"

	"github.com/claisne/lencha/controllers"
	"github.com/claisne/lencha/middlewares"
	"github.com/claisne/lencha/problems"

	"github.com/gorilla/mux"
)

// Returns the handlers from CreateRoutes,
// augmented with some middlewares
func GenerateHandlers() http.Handler {
	var handlers http.Handler

	routesHandlers := CreateRoutes()

	handlers = middlewares.LoggingRequest(routesHandlers)

	return handlers
}

// Returns all the hanlders, attached to their
// respectives routes
func CreateRoutes() http.Handler {
	// Router definition
	router := mux.NewRouter()
	router.StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()
	problemsApiRouter := apiRouter.PathPrefix("/problems").Subrouter()

	// Controllers
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/problems", controllers.Problems)
	router.HandleFunc("/problems/{problem:[A-Za-z]+}", controllers.Problem)
	router.HandleFunc("/profile", middlewares.RequireLogged(http.HandlerFunc(controllers.Profile))).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	// Api
	apiRouter.HandleFunc("/users", controllers.ApiUsers).Methods("GET")
	// problemsApiRouter.HandleFunc("/", controllers.ApiProblems).Methods("GET")

	// Problems Api
	problemsApiRouter.HandleFunc("/reverse", middlewares.RequireApiKey(problems.HandlerFromStateHandler(problems.Reverse)))
	problemsApiRouter.HandleFunc("/equation", middlewares.RequireApiKey(problems.HandlerFromStateHandler(problems.Equation)))
	problemsApiRouter.HandleFunc("/maze", middlewares.RequireApiKey(problems.HandlerFromStateHandler(problems.Maze)))

	// Static assets
	router.PathPrefix("/fonts").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./static/fonts/"))))
	router.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css/"))))
	router.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js/"))))

	// 404
	router.NotFoundHandler = http.HandlerFunc(notFound)

	return router
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/404.html")
}
