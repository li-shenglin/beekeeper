package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account  string `gorm:"unique;type:varchar(128);not null"`
	Name     string `gorm:"type:varchar(128);not null"`
	PassWord string `gorm:"type:varchar(256);not null"`
}
