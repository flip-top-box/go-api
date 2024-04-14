package main

import (
	"GOAPI/cmd/api"
	"GOAPI/config"
	"GOAPI/db"
	"database/sql"
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

	initStorage(dataBase)

	server := api.NewServer(host, dataBase)
	if err := server.Run(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func initStorage(dataBase *sql.DB) {
	err := dataBase.Ping()
	if err != nil {
		log.Fatal("Error connecting to DataBase: ", err)
	}

	log.Println("Database Connected Successfully...")
}
