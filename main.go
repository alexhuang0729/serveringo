package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
