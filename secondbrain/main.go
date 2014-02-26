package main

import (
	"database/sql"
	"net/http"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func SetupDb() *sql.DB {
	db, _ := sql.Open("postgres", "postgres://localhost/secondbrain?sslmode=disable")
	return db
}

func main() {

	m := routing()

	// start the application
	// Run calls ListenAndServe
	m.Run()
}

func routing() *martini.ClassicMartini {
	// create a martini mux with happy defaults
	m := martini.Classic()

	// add a middlware handler to enable rendering
	m.Use(render.Renderer(render.Options{Layout: "layout"}))

	// add a service to handle note storage
	m.MapTo(NewNoteStorage(SetupDb()), (*NoteStorer)(nil))

	// configure routes
	m.Get("/", ShowNotes)
	m.Post("/create", binding.Bind(Note{}), NewNote)

	return m
}

func ShowNotes(r render.Render, ns NoteStorer) {
	notes := ns.All()
	r.HTML(200, "books", notes)
}

func NewNote(note Note, w http.ResponseWriter, r *http.Request, ns NoteStorer) {
	ns.Add(note)
	http.Redirect(w, r, "/", http.StatusCreated)
}
