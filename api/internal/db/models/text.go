package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	NameID uint
	Name   Translation `gorm:"foreignKey:NameID`
}

type Translation struct {
	ID       uint   `gorm:"primaryKey"`
	Key      string `gorm:"not null"`
	Language string `gorm:"not null"`
	Text     string `gorm:"not null"`
}
