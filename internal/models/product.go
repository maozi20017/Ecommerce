package models

import "gorm.io/gorm"

type Product struct {
	Price       int
	Name        string
	Description string
	gorm.Model
}
