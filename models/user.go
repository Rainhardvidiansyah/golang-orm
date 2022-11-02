package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	UserName  string
	Email     string
	Pasword   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
