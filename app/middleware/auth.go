package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		admPass := os.Getenv("ADMIN_PASS")
		reqPass, ok := c.GetPostForm("password")
		if !ok || reqPass != admPass {
			c.AbortWithStatusJSON(200, "invalid request")
			return
		}

		c.Next()
	}
}
