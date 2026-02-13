// Package routes contains all the http routes
package routes

import "github.com/gin-gonic/gin"

func RegisterAdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}
}
