package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/controller"
)

func AddLandscape() gin.HandlerFunc {
	return func(c *gin.Context) {
		placeIDStr, ok := c.GetPostForm("place_id")
		if !ok {
			log.Println("place id missing")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing Place ID",
			})
			return
		}

		qualityStr, ok := c.GetPostForm("quality")
		if !ok {
			log.Println("quality is missing")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing quality value",
			})
			return
		}

		fileHeader, err := c.FormFile("image")
		if err != nil || fileHeader == nil {
			log.Println("no image(s) attached")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No image attached",
			})
			return
		}

		placeID, err := strconv.Atoi(placeIDStr)
		if err != nil {
			log.Println("str conv failed: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error with PlaceID",
			})
			return
		}

		quality, err := strconv.Atoi(qualityStr)
		if err != nil {
			log.Println("str conv failed: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error converting quality",
			})
			return
		}

		log.Println("saving image with place ID", placeID)
		landscape, cErr := controller.AddLandscape(uint(placeID), fileHeader, quality)
		if cErr != nil {
			log.Println("error occured while adding landscape: cleanup in progress: ", err)
			deletedBird, err := controller.RemoveLandscape(landscape.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "request was unsuccessful, cleanup failed too",
				})
				return
			}
			log.Println("cleanup completed..", deletedBird)

			log.Println("landscape image saving failed: ", cErr)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": cErr.Error(),
			})
			return
		}

		log.Println("landscape image saved..", landscape)
		c.JSON(200, "landscape image upload complete")
	}
}
