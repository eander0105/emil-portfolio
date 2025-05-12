package server

import "github.com/gin-gonic/gin"

type Healthcheck struct {
	Status bool `json:"status"`
}

func registerGeneralRoutes(router *gin.Engine) {
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, Healthcheck{Status: true})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Random test endpoint
	router.GET("/example/:name", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Query("age")
		c.JSON(200, gin.H{
			"message": "Hello " + name + " " + age,
		})
	})
}
