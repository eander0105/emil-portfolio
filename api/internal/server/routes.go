package server

import (
	"emil-portfolio/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Register general routes
	auth.RegisterRoutes(router.Group("/auth"))
	registerGeneralRoutes(router)

	return router
}
