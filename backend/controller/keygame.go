package controller

import (
	"net/http"

	"G13.com/backend/configs"
	"G13.com/backend/entity"
	"github.com/gin-gonic/gin"
)

func CreateKeyGame(c *gin.Context) {
	var keygame entity.KeyGame
	if err := c.ShouldBindJSON(&keygame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้าง record ใหม่ลง DB
	if err := configs.DB().Create(&keygame).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ส่งข้อมูล category ที่เพิ่งสร้างกลับไปให้ frontend
	c.JSON(http.StatusOK, keygame)
}
func FindKeyGame(c *gin.Context) {
	var keygame []entity.KeyGame
	if err := configs.DB().Preload("Game").Find(&keygame).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, keygame)
}

func DeleteKeyGameById(c *gin.Context) {
	id := c.Param("id")
	if tx := configs.DB().Exec("DELETE FROM KeyGame WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted succesful"})
}
