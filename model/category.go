package model

import "gorm.io/gorm"

type Category struct {
	Name string `gorm:"not null; size:255"`
	gorm.Model
}
