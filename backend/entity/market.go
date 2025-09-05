package entity

import (
	"gorm.io/gorm"
)

type Market struct {
	gorm.Model
	GameID uint
}
