package middlewares

import (
	"net/http"

	"github.com/claisne/lencha/models"
	"github.com/claisne/lencha/session"

	ctx "github.com/gorilla/context"
)

func RequireLogged(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := session.Store.Get(r, "lencha-session")

		if id, ok := session.Values["id"]; ok {
			u, err := models.GetUserById(id.(int))
			if err != nil {
				http.Redirect(w, r, "/", http.StatusFound)
				ctx.Clear(r)
				return
			} else {
				ctx.Set(r, "user", u)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
			ctx.Clear(r)
			return
		}

		handler.ServeHTTP(w, r)

		ctx.Clear(r)
	}
}
