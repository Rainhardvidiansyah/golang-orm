package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        int `gorm:"primaryKey:autoIncrement"`
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
