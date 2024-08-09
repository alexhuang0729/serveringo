package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainRoute)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
