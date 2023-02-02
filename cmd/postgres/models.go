package postgres

import (
	"gorm.io/gorm"
)

// ID, createdAt, updatedAt are automatically added by gorm
type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"uniqueIndex"`
	Age    uint8
	Active bool `gorm:"default:true"`
}
