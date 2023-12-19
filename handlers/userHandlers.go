package handlers

import (
	"log"
	"myGram/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Set CreatedAt and UpdatedAt timestamps
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = time.Now()

		// Create the user in the database
		if err := db.Create(&newUser).Error; err != nil {
			log.Println("Failed to create user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, newUser)
	}
}
