// Package server Starts and configs gin HTTP server
package server

import (
	"abhedya_2026/internal/configs"
	"abhedya_2026/internal/handlers"
	"abhedya_2026/internal/repository"
	"abhedya_2026/internal/routes"
	"abhedya_2026/internal/services"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	configs.LoadEnv()
	database := configs.ConnectDB()
	repo := repository.NewUserRepository(database)
	service := services.NewUserService(repo)
	handlers.NewUserHandler(service)

	g := gin.Default()
	routes.RegisterUserRoutes(g)
	routes.RegisterAdminRoutes(g)
	routes.RegisterAdminRoutes(g)

	return g
}
