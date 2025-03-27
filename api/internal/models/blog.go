package models

import "gorm.io/gorm"

type BlogTag struct {
	gorm.Model
	ID   int
	Name string
}

type Blog struct {
	gorm.Model
	ID      int
	Title   string
	Content string
	Tags    []BlogTag `gorm:"many2many:blog_tags;"`
}
