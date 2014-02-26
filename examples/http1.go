package main

import (
	"io"
	"net/http"
	"time"
)

// START OMIT
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Welcome to the Go Steel Programmers Meetup!")
	})
	http.HandleFunc("/show-time", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>"+time.Now().String()+"</h1>")
	})
	http.ListenAndServe(":8090", nil)
}

// END OMIT
