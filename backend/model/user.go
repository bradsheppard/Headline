package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id     int `gorm:"index"`
	Name   string
	Topics []*Topic `gorm:"many2many:user_topics;"`
}
