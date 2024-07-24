package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
