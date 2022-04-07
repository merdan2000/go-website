package main

import (
	_ "database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/merdan2000/go-website/pkg/handler"
)

func main() {
	addr := flag.String("addr", "4000", "адрис веб сервера")
	//dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Название PostgreSQL источника данных")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := handler.NewTemplateCache("ui/template/")
	app := handler.Application{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		TemplateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Router(),
	}

	infoLog.Printf("Запуск сервера на http://127.0.0.1:%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
