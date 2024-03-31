package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.createTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", s.getTasksHandler).Methods(http.MethodGet)
}

func (s *TasksService) createTasksHandler(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	defer req.Body.Close()
	fmt.Printf("%s", body)
}

func (s *TasksService) getTasksHandler(res http.ResponseWriter, req *http.Request) {

}
