package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentID string `valid:"required~required is student ID, matches(^[BMD]\\d{7}$)~invalid format studentID"`
	Email string `valid:"required,email~Invalid format Email"`
	Phone string `valid:"required~Phone is required, stringlength(10|10)~not format for phone"`
}