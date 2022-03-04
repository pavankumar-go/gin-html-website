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
func AdminAPIWildlifeUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "wildlife_upload.html", gin.H{
			"Host": Host,
		})
	}
}

func AdminAPIWildlifePlaceUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "wildlife_place_upload.html", gin.H{
			"Host": Host,
		})
	}
}

func AdminAPIWildlifePlacePatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "update_wildlife_place_thumbnail.html", gin.H{
			"Host": Host,
		})
	}
}



// Admin frontend
func AdminAPILandscapeUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "landscape_upload.html", gin.H{
			"Host": Host,
		})
	}
}

func AdminAPILandscapePlaceUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "landscape_place_upload.html", gin.H{
			"Host": Host,
		})
	}
}

func AdminAPILandscapePlacePatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "update_landscape_place_thumbnail.html", gin.H{
			"Host": Host,
		})
	}
}
