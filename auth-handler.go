package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		sendResponse(w, false, "Input data is not valid.")
		return
	}
	r.Body.Close()

	if user.Name == "" || user.Email == "" || user.Password == "" {
		sendResponse(w, false, "All fields are required.")
		return
	}

	user.Email = strings.ToLower(user.Email)

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		sendResponse(w, false, "Email address is not valid.")
		return
	}

	var count int
	gormDatabase.Model(&User{}).Where(&User{
		Email: user.Email,
	}).Limit(1).Count(&count)

	if count != 0 {
		sendResponse(w, false, "Email address is not available.")
		return
	}

	b, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	user.Password = string(b)
	user.UUID = uuid.NewV4().String()
	gormDatabase.Create(&user)

	sendResponse(w, true, "Sign up successful. Please check your email for further instructions.")
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "user-session")
	if err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		sendResponse(w, false, errPlatformUnstable)
		return
	}

	user.Email = strings.ToLower(user.Email)

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		log.Println(err)
		sendResponse(w, false, "Email address is not valid.")
		return
	}

	if user.Email == "" || user.Password == "" {
		sendResponse(w, false, "All fields are required.")
		return
	}

	var userX User
	gormDatabase.Where(&User{
		Email: user.Email,
	}).First(&userX)
	if userX.ID == 0 {
		sendResponse(w, false, "Email and/or password is not valid.")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userX.Password), []byte(user.Password)); err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			log.Println(err)
		}
		sendResponse(w, false, "Email and/or password is not valid.")
		return
	}

	session.Values["userID"] = userX.ID

	session.Options.HttpOnly = true

	session.Save(r, w)

	sendResponse(w, true, "Sign in successful.")
}

func signOutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "user-session")
	if err != nil {
		log.Println(err)
	}

	session.Options = &sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}

	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
