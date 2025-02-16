package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Appointment struct {
    ID        uuid.UUID    `gorm:"primaryKey"`
    Title     string
    Start     time.Time    `gorm:"type:timestamp"`
    End       time.Time    `gorm:"type:timestamp"`
    CreatorID  uuid.UUID 
    Creator    User
}