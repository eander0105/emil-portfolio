package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/register", registerHandler)
	r.POST("/login", loginHandler)
}
