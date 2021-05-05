package handlers

import (
	"net/http"

	"github.com/saravanakkumar/goWeb/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.ParseTemplate(w, "about.html")
}
