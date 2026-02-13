// Package models contains all models used in project
package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;unique"`
	Username  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email     string    `gorm:"type:varchar(150);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Role      string    `gorm:"type:varchar(50);default:user"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Submission struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	UserID          uint `gorm:"index;not null"`
	QuestionID      uint `gorm:"index;not null"`
	Score           int
	LastSubmittedAt time.Time `gorm:"autoUpdateTime"`
}

type Question struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Question       string    `gorm:"type:text;not null"`
	QuestionNumber int       `gorm:"uniqueIndex;not null"`
	Answer         string    `gorm:"type:text"`
	ImgSrc         string    `gorm:"type:varchar(255)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
