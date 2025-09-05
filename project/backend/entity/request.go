package entity

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	Reason    string
	Date      time.Time
	UserRefer uint `gorm:"not null;index:idx_user_game,unique"`
	GameRefer uint `gorm:"not null;index:idx_user_game,unique"`
}
