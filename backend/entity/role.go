package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description" gorm:"unique"`
	Users       []User `gorm:"foreignKey:RoleRefer"`
}
