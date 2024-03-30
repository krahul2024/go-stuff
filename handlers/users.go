package handlers

import (
	"log"
	"net/http"
)

type Users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

func (u *Users) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	u.l.Printf("Route called : %v\n", req.URL)

	res.Write([]byte("This is the get all user route!"))
}

func (u *Users) GetUser(res http.ResponseWriter, req *http.Request) {
	u.l.Printf("Route called : %v\n", req.URL)

	res.Write([]byte("This is the get user route!"))
}

func (u *Users) AddUser(res http.ResponseWriter, req *http.Request) {
	u.l.Printf("Route called : %v\n", req.URL)

	res.Write([]byte("This is the add users route!"))
}

func (u *Users) UpdateUser(res http.ResponseWriter, req *http.Request) {
	u.l.Printf("Route called : %v\n", req.URL)

	res.Write([]byte("This is the update user route!"))
}

func (u *Users) DeleteUser(res http.ResponseWriter, req *http.Request) {
	u.l.Printf("Route called : %v\n", req.URL)

	res.Write([]byte("This is the delete user route!"))
}
