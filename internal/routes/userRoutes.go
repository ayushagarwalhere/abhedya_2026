// Package routes contains all the http routes
package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}
}
