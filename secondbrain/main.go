package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/codegangsta/martini"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

// example env PSQL_URI=postgres://localhost/secondbrain?sslmode=disable"
func SetupDb() *sql.DB {
	db, _ := sql.Open("postgres", os.Getenv("PSQL_URI"))
	return db
}

func main() {
	m := Shaken()
	// start the application
	// Run calls ListenAndServe
	m.Run()
}

func Shaken() *martini.ClassicMartini {
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
	r.HTML(200, "notes", notes) // HLRENDER
}

func NewNote(note Note, r render.Render, ns NoteStorer) {
	ns.Add(note)
	notes := ns.All()
	r.HTML(http.StatusCreated, "notes", notes) // HLRENDER
}
