package main

import (
	"io"
	"net/http"
)

// START OMIT
func main() {
	s := &Server{}
	http.ListenAndServe(":8090", s)
}

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "(╯°□°）╯︵ ┻━┻")
}

// END OMIT
