package handlers

import (
	"log"
	"net/http"
)

type Index struct {
	l *log.Logger
}

func NewIndex(l *log.Logger) *Index {
	return &Index{l}
}

func (idx *Index) IndexRoute(res http.ResponseWriter, req *http.Request) {
	idx.l.Printf("Route called %v\n", req.URL)

	res.Write([]byte("Welcome to the index page!\n"))
}
