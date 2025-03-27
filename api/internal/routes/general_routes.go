package routes

import "github.com/gin-gonic/gin"

type Healthcheck struct {
	Status bool `json:"status"`
}

func RegisterGeneralRoutes(router *gin.Engine) {
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, Healthcheck{Status: true})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
