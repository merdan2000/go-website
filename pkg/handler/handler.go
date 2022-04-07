package handler

import (
	"net/http"
)

func (app *Application) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", 404)
		return
	}
	app.cookieClient(w, r, "")
	app.render(w, r, "home-page.html")

}

func (app *Application) authPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth" {
		http.Error(w, "not found", 404)
		return
	}
	app.authData(w, r)
	app.render(w, r, "auth-page.html")
}
