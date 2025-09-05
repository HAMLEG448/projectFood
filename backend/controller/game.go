package controller

import (
	"net/http"

	"G13.com/backend/configs"
	"G13.com/backend/entity"
	"github.com/gin-gonic/gin"
)

func FindGames(c *gin.Context) {
	var games []entity.Game
	if err := configs.DB().
		Preload("Categories").
		Find(&games).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}
func CreateGame(c *gin.Context) {
	var input struct {
		GameName        string `json:"game_name" binding:"required"`
		BasePrice       int    `json:"base_price" binding:"required"`
		AgeRating       int    `json:"age_rating"`
		ImgSrc          string `json:"img_src"`
		Minimum_spec_id int    `json:"minimum_spec_id" binding:"required"`
		CategoriesID    int    `json:"categories_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	games := entity.Game{
		GameName:       input.GameName,
		Baseprice:      input.BasePrice,
		AgeRating:      input.AgeRating,
		ImgSrc:         input.ImgSrc,
		Minimum_specID: uint(input.Minimum_spec_id),
		CategoriesID:   input.CategoriesID,
	}
	if err := configs.DB().Create(&games).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}
