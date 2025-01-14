package entity

import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	StudentID string `valid:"requried~Student ID is required, matches(^[BMD]\\d{7}$)"`
}