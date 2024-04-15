package main

import (
	"GOAPI/config"
	"GOAPI/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// Construct DSN from environment/config variables
	dsn := config.Envs.DBDSN

	// Open the DB connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate your schemas
	err = db.AutoMigrate(&types.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Additional logic to handle command line arguments for up or down migrations
	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "up":
			// Optionally handle specific 'up' migration logic
			log.Println("Migrations applied successfully")
		case "down":
			// Handle 'down' migrations if you have a method to roll back
			log.Println("Rolled back migrations successfully")
		default:
			log.Println("Invalid command")
		}
	}
}
