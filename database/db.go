package database

import (
	"fmt"
	"log"
	"tugas-dua/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "naufa123"
	dbname   = "tugas_kedua"
	db       *gorm.DB
	err      error
)

func GetConnection() *gorm.DB {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}
	log.Println("DB Connecting Established...")
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Order{})
	return db
}
