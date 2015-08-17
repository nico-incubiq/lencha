package middlewares

import (
	"net/http"

	"github.com/claisne/lencha/models"

	ctx "github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

func SetUserIdInSession(w http.ResponseWriter, r *http.Request, id int) {
	session, _ := SessionStore.Get(r, "lencha-session")
	session.Values["id"] = id
	session.Save(r, w)
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, "lencha-session")
	delete(session.Values, "id")
	session.Save(r, w)
}

func IsLogged(r *http.Request) bool {
	session, _ := SessionStore.Get(r, "lencha-session")
	_, exist := session.Values["id"]
	return exist
}

func RequireLogged(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := SessionStore.Get(r, "lencha-session")

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
