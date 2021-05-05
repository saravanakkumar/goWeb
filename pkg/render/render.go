package render

import (
	"html/template"
	"log"
	"net/http"
)

func ParseTemplate(w http.ResponseWriter, tmpl string) {
	out, err := template.ParseFiles("../../templates/" + tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = out.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
