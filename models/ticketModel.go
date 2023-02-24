package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID uint
	User   User
	Status bool
}
