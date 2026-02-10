// Package server Starts and configs gin HTTP server
package server

import (
	"log"

	"abhedya_2026/internal/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	g := gin.Default()
	routes.RegisterRoutes(g)
	err := g.Run(":3000")
	if err != nil {
		log.Fatal("server couldn't start")
	}
}
