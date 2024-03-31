package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MysqlStorage struct {
	db *sql.DB
}

func NewMysqlStorage(config mysql.Config) *MysqlStorage {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping() // this is for validation of the database and data source
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database!")
	return &MysqlStorage{db: db}
}

func (s *MysqlStorage) Init() (*sql.DB, error) {
	// three tables for projects, tasks and users
	return s.db, nil
}

func (s *MysqlStorage) createProjectsTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
			name VARCHAR(255) NOT NULL, 
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
			PRIMARY KEY(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8; 
	`)
	return err
}

func (s *MysqlStorage) createTasksTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
			name VARCHAR(255) NOT NULL, 
			status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO', 
			projectId INT UNSIGNED NOT NULL, 
			assignedToID INT UNSIGNED NOT NULL, 
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
			PRIMARY KEY (id), 
			FOREIGN KEY (assignedToID) REFERENCES users(id), 
			FOREIGN KEY (projectId) REFERENCES projects(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	return err
}

func (s *MysqlStorage) createUsersTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
			email VARCHAR(255) NOT NULL, 
			firstName VARCHAR(255) NOT NULL, 
			lastName VARCHAR(255) NOT NULL, 
			password VARCHAR(500) NOT NULL, 
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
			PRIMARY KEY(id), 
			UNIQUE KEY(email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8
	`)
	return err
}
