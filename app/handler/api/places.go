package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/controller"
)

func AddWildlifePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		placeName, ok := c.GetPostForm("name")
		if !ok {
			log.Println("place name missing")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing Place Name",
			})
			return
		}

		// NOTE: an handler for each place should be of this shortname
		// ex: shortname = blr (bangalore)
		// api should be '/places/blr'
		shortName, ok := c.GetPostForm("shortname")
		if !ok {
			log.Println("place name missing")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing Place Name",
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

		log.Println("creating wildlife place and saving background image for: ", placeName)
		place, cErr := controller.AddWildlifePlace(placeName, shortName, fileHeader)
		if cErr != nil {
			log.Println("error occured cleanup in progress: ", err)
			deletedBird, err := controller.RemoveWildlifePlace(place.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "request was unsuccessful, cleanup failed too",
				})
				return
			}
			log.Println("cleanup completed..", deletedBird)

			log.Println("adding wildlife place and saving background image had failed: ", cErr)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create wildlife place: " + cErr.Error(),
			})
			return
		}

		log.Printf("wildlife place created with ID: %d and background is saved..", place.ID)
		c.JSON(200, "wildlife place created and background image upload complete")
	}
}

func DeleteWildlifePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		placeIDStr := c.Param("id")
		placeID, err := strconv.Atoi(placeIDStr)
		if err != nil {
			log.Println("failed to strconv.Atoi : ", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "invalid id",
			})
			return
		}

		log.Println("deleting place: ", placeID)
		ok, err := controller.RemoveWildlifePlace(uint(placeID))
		if err != nil || !ok {
			log.Println("error deleting wildlife place: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "failed to delete wildlife place: " + err.Error(),
			})
			return
		}
		c.JSON(200, "wildlife place deleted successfully..")
	}
}
