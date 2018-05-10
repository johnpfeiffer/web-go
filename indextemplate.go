package main

import "html/template"

// NoData is an empty struct as I do not pass anything into the template
type NoData struct{}

// GetIndexTemplate returns the index.html template https://golang.org/pkg/html/template/#Template
func GetIndexTemplate() *template.Template {
	indexTemplate := template.Must(template.ParseFiles("base.tmpl", "index.html"))
	return indexTemplate
}
