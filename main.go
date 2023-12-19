package main

import (
	"fmt"
	"log"
	"myGram/config"
	"myGram/handlers"
	"myGram/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := config.NewPostgresConfig()

	r := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	log.Println("Database connection successful")

	r.POST("/users", handlers.CreateUserHandler(db))

	r.GET("/photos", handlers.GetAllPhotos(db))
	r.GET("/comments", handlers.GetAllComments(db))
	r.GET("/socialmedia", handlers.GetAllSocialMedia(db))

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}

}
