package infrastructure

import (
	"log"
	"os"

	"favorite-book/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:password@localhost:5432/db_book"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.Book{}); err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal(err)
	}

	return db
}
