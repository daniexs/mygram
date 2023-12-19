package handlers

import (
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllComments(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var comments []models.Comment
		if err := db.Find(&comments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
			return
		}

		c.JSON(http.StatusOK, comments)
	}
}
