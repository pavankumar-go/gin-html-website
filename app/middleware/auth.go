package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		admPass := os.Getenv("ADMIN_PASS")
		reqPass, ok := c.GetPostForm("password")
		if !ok || reqPass != admPass {
			log.Println("auth failed")
			c.AbortWithStatusJSON(200, "invalid request")
			return
		}

		c.Next()
	}
}

func init() {
	err := godotenv.Load(os.Getenv("SECRET_PATH"))
	if err != nil {
		log.Fatalln("failed to load env: ", err)
	}
}
