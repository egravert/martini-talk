package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { // HL
		io.WriteString(w, "Welcome to the Go Steel Programmers Meetup!") // HL
	}) // HL
	http.ListenAndServe(":8090", nil) // HL
}
