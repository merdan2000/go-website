package handler

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, name string) {
	ts, ok := app.TemplateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("шаблон %s не существует", name))
		return
	}
	err := ts.Execute(w, "")
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *Application) cookieClient(w http.ResponseWriter, r *http.Request, name string) {
	cookei, err := r.Cookie("enter")
	if err != nil {
		cookei = &http.Cookie{
			Name:     "enter",
			Value:    "?????",
			HttpOnly: true,
			Secure:   false,
		}
	}

	if cookei.Secure == true {
		if name != "auth-page.html" {
			app.render(w, r, name)
		}
	}

}
