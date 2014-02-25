package main

import (
	"fmt"
	"github.com/codegangsta/martini"
)

type User struct {
	Name string
}

func main() {
	m := martini.Classic()
	m.Use(authorize)
	m.Use(logUser)
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}

func authorize(c martini.Context) {
	c.Map(User{"Eric"})
}

func logUser(u User) {
	fmt.Println(u.Name)
}
