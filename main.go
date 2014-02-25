package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := routing()
	m.Use(render.Renderer(render.Options{Layout: "layout"}))
	m.Run()
}

func routing() *martini.ClassicMartini {
	m := martini.Classic()
	m.Get("/", func(r render.Render) { r.HTML(200, "home", nil) })
	m.Get("/books", ShowBooks)
	m.Get("/books/new", NewBook)
	m.Post("/books/new", CreateBook)

	return m
}

func ShowBooks(r render.Render) {
	r.HTML(200, "books", nil)
}

func NewBook(r render.Render) {
	r.HTML(200, "books", nil)
}


func CreateBook(r render.Render) {

}
