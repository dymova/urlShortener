package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlShortener/auth"
)

func JwtAuth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, err := auth.VerifyJWTToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
