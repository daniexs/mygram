package handlers

import (
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPhotos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var photos []models.Photo
		if err := db.Find(&photos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
			return
		}

		c.JSON(http.StatusOK, photos)
	}
}
