package handlers

import (
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllSocialMedia(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var socialMedia []models.SocialMedia
		if err := db.Find(&socialMedia).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch social media data"})
			return
		}

		c.JSON(http.StatusOK, socialMedia)
	}
}
