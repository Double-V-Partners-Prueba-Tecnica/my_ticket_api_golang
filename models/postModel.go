package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
}

type Ticket struct {
	gorm.Model
	UserID uint
	User   User
	Status string
}
