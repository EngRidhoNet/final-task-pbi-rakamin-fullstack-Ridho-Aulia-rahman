package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"not null"`
    Email     string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
