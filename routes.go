package main

import (
	"net/http"
)

func serveRoutes() {
	http.HandleFunc("/", mainRoute)
}
