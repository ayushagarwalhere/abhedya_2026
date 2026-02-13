// Package repository contains all the dbQueries
package repository

import (
	"abhedya_2026/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}
