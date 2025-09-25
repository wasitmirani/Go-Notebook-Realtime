package database

import (
	"GoNotebookRealtime/config"
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB(cfg config.Config) *gorm.DB {
	// Build the connection string
	fmt.Printf("DB Config: Host=%s, Port=%s, User=%s, DBName=%s\n", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	fmt.Println("Connecting to database with DSN:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
