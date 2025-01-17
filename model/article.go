package model

import "gorm.io/gorm"

type Status int

const (
	StatusDraft   Status = 0 // 草稿
	StatusPublish Status = 1 // 发布
)

type Article struct {
	Title      string     `gorm:"not null; size:255"`
	SubTitle   string     `gorm:"null; size:255"`
	Content    string     `gorm:"not null;"`
	Categories []Category `gorm:"many2many:article_categories;"`
	Tags       []Tag      `gorm:"many2many:article_tags;"`
	Status     Status     `gorm:"not null;"`
	gorm.Model
}
