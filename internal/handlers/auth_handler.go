// Package handlers contains all the HTTP layer
package handlers

import (
	"net/http"

	"abhedya_2026/internal/models"
	"abhedya_2026/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userServices *services.UserServices
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserHandler(s *services.UserServices) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var req SignupRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"message": "encountered some error",
			"error":   err,
		})
		return
	}
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.userServices.Signup(&user); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "user created successfully",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"message": "encountered some error",
			"error":   err,
		})
		return
	}
	if req.Username == "" || req.Password == "" {
		c.JSON(501, gin.H{
			"error": "Username or Password is empty",
		})
		return
	}
	token, err := h.userServices.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(501, gin.H{
			"error": err,
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authToken", token, 60*60*48, "/", "", true, true)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
	})
}
