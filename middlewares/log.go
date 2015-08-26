package middlewares

import (
	"net"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func LoggingRequest(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log only if not haproxy
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if ip != "127.0.0.1" {
			log.WithFields(log.Fields{"method": r.Method, "url": r.URL.String(), "remote_addr": r.RemoteAddr}).Info("HTTP request")
		}

		handler.ServeHTTP(w, r)
	}
}
