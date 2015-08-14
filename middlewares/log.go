package middlewares

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func LoggingRequest(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		log.WithFields(log.Fields{"method": r.Method, "url": r.URL.String()}).Info("HTTP request")
	}
}
