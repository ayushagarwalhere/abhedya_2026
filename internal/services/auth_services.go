// Package services contains all the application logic
package services

import (
	"log"

	"abhedya_2026/internal/models"
	"abhedya_2026/internal/repository"
	"abhedya_2026/internal/utils"
)

type UserServices struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserServices {
	return &UserServices{r}
}

func (s *UserServices) Signup(user *models.User) error {
	user.Role = "USER"
	hashedPassword, err := utils.HashString(user.Password)
	if err != nil {
		log.Println("encountered some error")
		return err
	}
	user.Password = string(hashedPassword)

	return s.repo.CreateUser(user)
}

