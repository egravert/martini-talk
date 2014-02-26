package main

import (
	"time"

	"github.com/codegangsta/martini"
)

// START OMIT
func main() {
	m := martini.Classic()     // HLCLASSIC
	m.Get("/", func() string { // HLROUTE
		return "Hello world!" // HLROUTE
	}) // HLROUTE
	m.Post("/create", func() string { // HLROUTE
		return "Thanks for your feedback" // HLROUTE
	}) // HLROUTE
	m.Any("/time", func() string { // HLROUTE
		return "The time is: " + time.Now().Format(time.Kitchen) // HLROUTE
	}) // HLROUTE
	m.Run()
}

// END OMIT
