package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
}

type List struct {
	gorm.Model
}
