package session

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("something-very-secret"))

func SetUserId(w http.ResponseWriter, r *http.Request, id int) {
	session, _ := Store.Get(r, "lencha-session")
	session.Values["id"] = id
	session.Save(r, w)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "lencha-session")
	delete(session.Values, "id")
	session.Save(r, w)
}

func IsLogged(r *http.Request) bool {
	session, _ := Store.Get(r, "lencha-session")
	_, exist := session.Values["id"]
	return exist
}
func GetId(r *http.Request) (int, error) {
	session, _ := Store.Get(r, "lencha-session")
	val, exist := session.Values["id"]
	if !exist {
		return -1, errors.New("Not logged in !")
	}

	if id, ok := val.(int); ok {
		return id, nil
	}

	return -1, errors.New("Could not retrieve the session data")
}
