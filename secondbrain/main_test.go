package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// Expect valudates expected vs actual and returns a message indicating type and value difference if
// they do not match
func expect(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

// Creates a new request. If content type is not empty, the Content-Type header will also be set
func NewRequest(method, urlStr, contentType, body string) *http.Request {
	request, _ := http.NewRequest(method, urlStr, strings.NewReader(body))

	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
	}
	return request
}

// Validates that the index page is accessible
// receives back a bunch of html...
func TestShowingNotes(t *testing.T) {

	m := Shaken() // a method on main that configures, but doesnt `Run` the martini instance
	response := httptest.NewRecorder()
	request := NewRequest("GET", "/", "", "")

	m.ServeHTTP(response, request)

	expect(t, http.StatusOK, response.Code)
}

// Creates a new note and validates that it exists in the db
func TestCreatingNotes(t *testing.T) {
	var startRows, endRows int

	db := SetupDb()

	defer db.QueryRow("DELETE FROM notes")

	db.QueryRow("SELECT COUNT(*) FROM notes").Scan(&startRows)
	m := Shaken() // a method on main that configures, but doesnt `Run` the martini instance

	response := httptest.NewRecorder()
	request := NewRequest("POST", "/create", "application/x-www-form-urlencoded; param=value", "note=always%20test%20your%20app&tags=")
	m.ServeHTTP(response, request)

	// validate http status
	expect(t, http.StatusCreated, response.Code)

	// validate row counts in db
	db.QueryRow("SELECT COUNT(*) FROM notes").Scan(&endRows)
	expect(t, startRows+1, endRows)
}
