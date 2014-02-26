package main

import (
	"net/http"
	"time"

	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic() // HLCLASSIC
	// START OMIT
	m.Any("/time", func() string { // HLROUTE
		return "The time is: " + time.Now().Format(time.Kitchen)
	})

	m.Post("/create", func() (int, string) { // HLROUTE
		return http.StatusCreated, "Thanks for your feedback"
	})

	m.Get("/hello/:name", func(p martini.Params) string { // HLROUTE
		return "Hello, " + p["name"]
	})
	// END OMIT
	m.Run()
}
