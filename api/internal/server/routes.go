package server

import (
	"github.com/eander0105/emil-portfolio/api/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Register general routes
	auth.RegisterRoutes(router.Group("/auth"))
	registerGeneralRoutes(router)

	return router
}
