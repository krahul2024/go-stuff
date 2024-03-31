package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()

	// tasks route
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(router)

	log.Printf("Starting the server at PORT %v", s.addr)
	err := http.ListenAndServe(s.addr, subRouter)
	if err != nil {
		log.Fatal("There was an error starting the server at port ", s.addr, "\n", err)
	}
}
