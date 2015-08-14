package controllers

import (
	"clem/lencha/middlewares"
	"clem/lencha/models"
	"clem/lencha/utils"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"

	"github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")

	user, err := models.GetUserByUsername(username)
	if err != nil {
		JSONResponse(w, models.Response{Success: false, Message: "No user with this username."}, http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		JSONResponse(w, models.Response{Success: false, Message: "Wrong password", Data: user}, http.StatusUnauthorized)
		return
	}

	middlewares.SetUserIdInSession(w, r, user.Id)
	JSONResponse(w, models.Response{Success: true}, http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	middlewares.DeleteSession(w, r)
	http.Redirect(w, r, "/", 302)
}

func GenerateApiKey() string {
	k := make([]byte, 32)
	io.ReadFull(rand.Reader, k)
	return fmt.Sprintf("%x", k)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	emailUpdate := r.FormValue("email-update") == "true"

	if !utils.IsAlphaNumeric(username) {
		JSONResponse(w, models.Response{Success: false, Message: "Your username must be alphanumeric"}, http.StatusUnauthorized)
		return
	}

	if len(username) < 3 || len(username) > 16 {
		JSONResponse(w, models.Response{Success: false, Message: "Your username must have a length between 3 and 16 characters."}, http.StatusUnauthorized)
		return
	}

	if len(password) < 3 || len(password) > 16 {
		JSONResponse(w, models.Response{Success: false, Message: "Your password must have a length between 3 and 16 characters."}, http.StatusUnauthorized)
		return
	}

	if len(email) != 0 && !utils.IsEmail(email) {
		JSONResponse(w, models.Response{Success: false, Message: "Your email is not valid"}, http.StatusUnauthorized)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		JSONResponse(w, models.Response{Success: false, Message: "Error creating the account. Try again please"}, http.StatusUnauthorized)
		return
	}

	user := models.User{Username: username, Hash: string(hash), Email: email, ApiKey: GenerateApiKey(), EmailUpdate: emailUpdate}

	err = models.CreateUser(&user)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			if err.Constraint == "username_unique" {
				JSONResponse(w, models.Response{Success: false, Message: "This username already exists."}, http.StatusUnauthorized)
				return
			}
			if err.Constraint == "email_unique" {
				JSONResponse(w, models.Response{Success: false, Message: "This email already exists."}, http.StatusUnauthorized)
				return
			}
		}
		// We dont handle this error
		JSONResponse(w, models.Response{Success: false, Message: "Error creating the account."}, http.StatusUnauthorized)
		return
	}

	JSONResponse(w, models.Response{Success: true, Message: "Your account is now created ! Sign in and you can start solving challenges.", Data: user}, http.StatusOK)
}
