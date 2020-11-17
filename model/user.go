package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID uuid.UUID		`gorm:"primaryKey;size:20;not null"`
	Handle string		`gorm:"size:15;index;not null"`
	Name string			`gorm:"size:50;not null"`
	Email string		`gorm:"unique;not null"`
	Password string		`gorm:"size:66;not null"`
}