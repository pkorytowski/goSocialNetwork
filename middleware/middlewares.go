package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"socialNetwork/token"
	"strconv"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func TheSameIdMiddlewareForPostOperation() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, _ := token.ExtractTokenID(c)
		mapbody := make(map[string]string)
		strbody, _ := io.ReadAll(c.Request.Body)
		json.Unmarshal(strbody, &mapbody)
		idx, _ := strconv.Atoi(mapbody["id"])

		if userId != uint(idx) {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
