package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// NewMySQLStorage initializes a new database connection using GORM.
func NewMySQLStorage(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
		return nil, err
	}
	return db, nil
}
