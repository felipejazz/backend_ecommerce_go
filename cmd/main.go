package main

import (
	"log"

	"github.com/felipejazz/ecommerce_go/cmd/api"
	"github.com/felipejazz/ecommerce_go/cmd/config"
	"github.com/felipejazz/ecommerce_go/cmd/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{

		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
