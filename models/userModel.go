package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string        `gorm:"not null"`
	Email        string        `gorm:"not null"`
	Password     string        `gorm:"not null"`
	Age          int           `gorm:"not null"`
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `gorm:"autoUpdateTime"`
	Photos       []Photo       `gorm:"foreignKey:UserID"`
	Comments     []Comment     `gorm:"foreignKey:UserID"`
	SocialMedias []SocialMedia `gorm:"foreignKey:UserID"`
}

type Photo struct {
	gorm.Model
	Title     string    `gorm:"not null"`
	Caption   string    `gorm:"not null"`
	PhotoURL  string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type PhotoComment struct {
	gorm.Model
	PhotoID   uint
	CommentID uint
}

type Comment struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	PhotoID   uint      `gorm:"not null"`
	Message   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type SocialMedia struct {
	gorm.Model
	Name           string    `gorm:"not null"`
	SocialMediaURL string    `gorm:"not null"`
	UserID         uint      `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
