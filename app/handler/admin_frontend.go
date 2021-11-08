package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var Host string

func init() {
	var ok bool
	Host, ok = os.LookupEnv("HOST_ADDRESS")
	if !ok {
		log.Fatalln("HOST_ADDRESS is unset")
	}
}

// Admin frontend
func AdminAPIBirdUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "bird_upload.html", gin.H{
			"Host": Host,
		})
	}
}

func AdminAPIPlaceUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "place_upload.html", gin.H{
			"Host": Host,
		})
	}
}

func AdminAPIPlacePatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "update_place_thumbnail.html", gin.H{
			"Host": Host,
		})
	}
}
