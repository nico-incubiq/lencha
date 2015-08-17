package middlewares

import (
	"net/http"

	"github.com/claisne/lencha/models"

	ctx "github.com/gorilla/context"
)

func RequireApiKey(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header["Api-Key"]
		if len(apiKeyHeader) != 1 {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("No API-KEY header!"))
			ctx.Clear(r)
			return
		}

		apiKey := apiKeyHeader[0]
		user, err := models.GetUserByApiKey(apiKey)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Api Key is incorrect!"))
			ctx.Clear(r)
			return
		}

		ctx.Set(r, "user", user)
		handler.ServeHTTP(w, r)
		ctx.Clear(r)
	}
}
