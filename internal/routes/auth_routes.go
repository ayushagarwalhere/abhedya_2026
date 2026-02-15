// Package routes contains all the http routes
package routes

import (
	"abhedya_2026/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, h *handlers.UserHandler) {
	auth := r.Group("/")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.Login)
	}
}
