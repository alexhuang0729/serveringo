package main

import (
	"fmt"
	"net/http"
)

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
