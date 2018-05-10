package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// HexColors wraps a list of colors as hexadecimal strings
type HexColors struct {
	Colors []string
}

// GetHexTemplate returns the parsed file as a template object
func GetHexTemplate() *template.Template {
	return template.Must(template.ParseFiles("base.tmpl", "components.tmpl", "hexcolorstemplate.html"))
}

func hexController(w http.ResponseWriter, r *http.Request) {
	colors := []string{}
	for i := 255; i <= 16711680; i = i * 256 {
		s := fmt.Sprintf("%06X", i)
		colors = append(colors, s)
	}
	data := HexColors{colors}
	hexTemplate.Execute(w, data)
}
