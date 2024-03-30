package main

import (
	"crud-mux/handlers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	l := log.New(os.Stdout, "\nCRUD-API\n", log.LstdFlags)
	godotenv.Load(".env")
	PORT, DATABASE_URL := os.Getenv("PORT"), os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatal("There was an error connecting to the database!")
	}
	defer db.Close()

	// handlers
	indexHandler := handlers.NewIndex(l)
	usersHandler := handlers.NewUsers(l)

	router := mux.NewRouter()
	getRouter := router.Methods(http.MethodGet).Subrouter()
	postRouter := router.Methods(http.MethodPost).Subrouter()
	putRouter := router.Methods(http.MethodPut).Subrouter()
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()

	// GET Routes
	getRouter.HandleFunc("/", indexHandler.IndexRoute)
	getRouter.HandleFunc("/users", usersHandler.GetAllUsers)
	getRouter.HandleFunc("/users/{id}", usersHandler.GetUser)

	// POST Routes
	postRouter.HandleFunc("/users", usersHandler.AddUser)

	// PUT Routes
	putRouter.HandleFunc("/users/{id}", usersHandler.UpdateUser)

	// delete Routes
	deleteRouter.HandleFunc("/users/{id}", usersHandler.DeleteUser)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", PORT),
		Handler: router,
	}

	log.Printf("Starting the server on PORT : %v", PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("There was an error starting the server!\n", err)
	}
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
