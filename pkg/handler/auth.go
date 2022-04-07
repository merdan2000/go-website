package handler

import (
	"fmt"
	"net/http"
)

type logInUser struct {
	Emaile   string
	Password string
}
type singUpUser struct {
	Username string
	Emaile   string
	Password string
}

func (app *Application) authData(w http.ResponseWriter, r *http.Request) {
	login := &logInUser{
		Emaile:   r.FormValue("email"),
		Password: r.FormValue("pswd"),
	}
	singUp := singUpUser{
		Username: r.FormValue("singIntxt"),
		Emaile:   r.FormValue("singInEmail"),
		Password: r.FormValue("singInpswd"),
	}
	if login.Emaile != "" && login.Password != "" {
		fmt.Println("ok login")
		authsetCookie(w, r, login.Password)
	}
	if singUp.Emaile != "" && singUp.Password != "" && singUp.Username != "" {
		fmt.Println("ok singin")
		authsetCookie(w, r, singUp.Password)
	}

}
func authsetCookie(w http.ResponseWriter, r *http.Request, value string) {

	fmt.Println("ok login")
	cookei, err := r.Cookie("enter")
	if err != nil {
		cookei = &http.Cookie{
			Name:     "enter",
			Value:    "?????",
			HttpOnly: true,
			Secure:   false,
		}
	}
	cookei = &http.Cookie{
		Name:     "enter",
		Value:    value,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, cookei)
	http.Redirect(w, r, "/", 302)

}
