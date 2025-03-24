package main

import "github.com/gin-gonic/gin"

type Healthcheck struct {
	Status bool `json:"status"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	// Healthcheck route
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, Healthcheck{Status: true})
	})

	err := router.Run(":3000")

	if err != nil {
		panic(err)
	}

}
