package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title         string `gorm:"not null"`
	DescriptionID uint
	Description   Translation `gorm:"foreignKey:DescriptionID"`
	BlogPosts     []BlogPost  `gorm:"many2many:project_blog_posts;constraints:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GithubURL     string
}
