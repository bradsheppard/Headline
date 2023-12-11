package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID     string
	Name   string
	Topics []*Topic `gorm:"many2many:user_topics;"`
}
