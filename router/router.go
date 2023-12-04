package router

import "github.com/gin-gonic/gin"

func Init() {
	// Initialize Router
	r := gin.Default()

	// Initialize Routes
	initializeRoutes(r)

	// Start Up Server
	r.Run(":8080")
}