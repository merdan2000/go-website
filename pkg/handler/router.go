package handler

import (
	"log"
	"net/http"
)

func (app Application) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homePage)
	mux.HandleFunc("/auth", app.authPage)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4001")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":4001", mux)
	log.Println(err)
	return mux
}
