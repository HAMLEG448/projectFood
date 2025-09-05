package controller

import (
	"net/http"

	"G13.com/backend/configs"
	"G13.com/backend/entity"
	"github.com/gin-gonic/gin"
)

func FindCategories(c *gin.Context) {
	var categories []entity.Categories
	if err := configs.DB().Raw("SELECT * FROM categories").Find(&categories).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
