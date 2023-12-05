package router

import (
	"github.com/JulioZittei/goportunities/config"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func Init() {
	logger = config.GetLogger("router")

	// Initialize Router
	r := gin.Default()
	r.Use(getLocaleContext)

	// Initialize Routes
	initializeRoutes(r)

	// Start Up Server
	r.Run(":8080")
}