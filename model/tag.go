package model

import "gorm.io/gorm"

type Tag struct {
	Name string `gorm:"not null; size:255"`
	gorm.Model
}
