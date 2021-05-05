package main

import (
	"net/http"

	"github.com/saravanakkumar/goWeb/pkg/handlers"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(port, nil)
}
