package handler

import (
	"html/template"
	"log"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	//snippets      *mysql.SnippetModel
	TemplateCache map[string]*template.Template
}
