package models

import "github.com/gofrs/uuid"

type User struct {
    ID              uuid.UUID `gorm:"primaryKey"`
    Name            string
    Username        string `gorm:"unique"`
    PreferredTimezone string
}