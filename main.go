package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	PORT, config := 3300, mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := NewMysqlStorage(config)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	apiServer := NewAPIServer(fmt.Sprintf(":%v", PORT), store)
	apiServer.Serve()
}
