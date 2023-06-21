package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"socialNetwork/token"
	"strconv"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.ValidateToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func TheSameIdMiddlewareForGetOperation() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenId, _ := token.ExtractTokenID(c)
		idstr, _ := c.Params.Get("id")
		id, _ := strconv.Atoi(idstr)
		if tokenId != uint(id) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
