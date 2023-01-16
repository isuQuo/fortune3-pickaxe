package controllers

import (
	"fmt"
	"fortune3-pickaxe/models"
	"net/http"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserService.Create(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Printf("User created: %+v\n", user)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)
}

func (u Users) Authenticate(w http.ResponseWriter, r *http.Request) {
	user, err := u.UserService.Authenticate(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/",  // This cookie will be available on all routes
		HttpOnly: true, // This cookie will not be available to JavaScript
	}
	http.SetCookie(w, &cookie)
	fmt.Printf("User signed in: %+v\n", user)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		return
	}
	fmt.Fprintf(w, "Current logged in user: %+v\n", cookie.Value)
}
