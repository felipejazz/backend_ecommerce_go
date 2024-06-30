package main

import (
	"database/sql"
	"log"

	"github.com/felipejazz/ecommerce_go/cmd/api"
	"github.com/felipejazz/ecommerce_go/config"
	"github.com/felipejazz/ecommerce_go/db"
	"github.com/go-sql-driver/mysql"
)

func main() {


	initStorage(db)

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}

m, err := migrate.NewWit
