package models

import (
    "time"
    "gorm.io/gorm"
)

type Photo struct {
    ID        uint      `gorm:"primaryKey"`
    Title     string    `gorm:"not null"`
    Caption   string
    PhotoUrl  string    `gorm:"not null"`
    UserID    uint      `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    User      User
}
