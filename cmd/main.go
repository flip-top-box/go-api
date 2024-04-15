package main

import (
	"GOAPI/cmd/api"
	"GOAPI/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := config.Envs
	host := cfg.Address
	dataBase, err := gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	initStorage(dataBase)

	server := api.NewServer(host, dataBase)
	if err := server.Run(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func initStorage(db *gorm.DB) {

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get underlying database from GORM: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error connecting to DataBase: ", err)
	}

	log.Println("Database Connected Successfully...")
}
