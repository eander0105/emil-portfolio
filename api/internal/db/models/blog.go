package models

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model
	Title   string
	Content string
	Tags    []Category `gorm:"many2many:blog_tags;constraints:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
