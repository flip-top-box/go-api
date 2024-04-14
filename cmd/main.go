package main

import (
	"GOAPI/cmd/api"
	"GOAPI/config"
	"GOAPI/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	cfg := config.Envs
	host := cfg.Address
	dataBase, err := db.NewMySQLStorage(mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPassword,
		Net:    "tcp",
		Addr:   cfg.DBAddress,
		DBName: cfg.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(host, dataBase)
	if err := server.Run(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
