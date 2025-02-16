package db

import (
	"fmt"
	"os"

	"github.com/gofrs/uuid"
	"github.com/rfaturriza/appointment-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func InitDB() {
    dsn := dsn()

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    DB = db

    db.Migrator().CreateTable(&models.User{})
    db.Migrator().CreateTable(&models.Appointment{})

    // seed data
    db.Create(&models.User{
        ID:       uuid.Must(uuid.NewV4()),
        Username: "admin",
        Name:     "Admin",
        PreferredTimezone: "Asia/Jakarta",
    })

    fmt.Println("Database connected successfully")
}

func dsn() string {
    dbUrl := os.Getenv("DATABASE_URL")
    if dbUrl != "" {
        return dbUrl
    }
    return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_SSL_MODE"),
        os.Getenv("DB_TIMEZONE"))
}