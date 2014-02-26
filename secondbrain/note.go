package main

import (
	"database/sql"
	"time"
)

type NoteStorer interface {
	All() []Note
	Add(note Note)
}

type Note struct {
	Id        int
	Note      string `form:"note"`
	Tags      string `form:"tags"`
	CreatedAt time.Time
}

type NoteStorage struct {
	db *sql.DB
}

func NewNoteStorage(db *sql.DB) *NoteStorage {
	return &NoteStorage{db}
}

func (ns *NoteStorage) All() []Note {
	rows, _ := ns.db.Query("SELECT id, tags, note, created_at FROM notes ORDER BY created_at desc")
	notes := []Note{}
	for rows.Next() {
		note := Note{}
		rows.Scan(&note.Id, &note.Tags, &note.Note, &note.CreatedAt)
		notes = append(notes, note)
	}
	return notes
}

func (ns *NoteStorage) Add(note Note) {
	ns.db.Query("INSERT INTO notes(note,tags) VALUES($1, $2)", note.Note, note.Tags)
}
