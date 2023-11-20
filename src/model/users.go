package model

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	gorm.Model
}
