package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id     string `gorm:"index"`
	Name   string
	Topics []*Topic `gorm:"many2many:user_topics;"`
}
