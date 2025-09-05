package entity

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	GameName       string     `json:"game_name" gorm:"type:varchar(512)"`
	KeyGameID      uint       `json:"key_id"`
	CategoriesID   int        `json:"categories_id"`
	Categories     Categories `json:"categories" gorm:"foreignkey:CategoriesID"`
	Date           time.Time  `json:"release_date" gorm:"autoCreateTime"`
	Baseprice      int        `json:"base_price"`
	ImgSrc         string     `json:"img_src"    gorm:"type:varchar(512)"`
	AgeRating      int        `json:"age_rating"`
	Status         string     `json:"status" gorm:"type:varchar(20);default:'pending';not null;index"`
	Minimum_specID uint       `json:"minimum_spec_id"`
	Requests       []Request  `gorm:"foreignKey:GameRefer"`
	Market         Market     `json:"market"`
}

// Hook function ไว้หลังสร้างเกมเสร็จแล้ว keygame จะเจนเอง
func (g *Game) AfterCreate(tx *gorm.DB) (err error) {
	kg := KeyGame{}
	if err = tx.Create(&kg).Error; err != nil {
		return err
	}
	return tx.Model(g).Update("key_game_id", kg.ID).Error
}
