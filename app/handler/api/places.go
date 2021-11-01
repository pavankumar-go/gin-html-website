package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/controller"
)

func AddPlace() gin.HandlerFunc {
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

		log.Println("creating place and saving background image for: ", placeName)
		place, cErr := controller.AddPlace(placeName, shortName, fileHeader)
		if cErr != nil {
			log.Println("error occured cleanup in progress: ", err)
			deletedBird, err := controller.RemovePlace(place.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "request was unsuccessful, cleanup failed too",
				})
				return
			}
			log.Println("cleanup completed..", deletedBird)

			log.Println("adding place and saving background image had failed: ", cErr)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create place: " + cErr.Error(),
			})
			return
		}

		log.Printf("place created with ID: %d and background is saved..", place.ID)
		c.JSON(200, "place created and background image upload complete")
	}
}

func DeletePlace() gin.HandlerFunc {
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
		ok, err := controller.RemovePlace(uint(placeID))
		if err != nil || !ok {
			log.Println("error deleting place: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "failed to delete place: " + err.Error(),
			})
			return
		}
		c.JSON(200, "place deleted successfully..")
	}
}

func UpdatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {

		placeIDStr, _ok := c.GetPostForm("id")
		placeID, err := strconv.Atoi(placeIDStr)
		if err != nil || !_ok {
			log.Println("failed to strconv.Atoi : ", err)
			c.AbortWithStatusJSON(400, gin.H{
				"message": "invalid id",
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

		log.Println("updating place and saving new background image for place ID: ", placeID)
		ok, err := controller.UpdatePlace(uint(placeID), fileHeader)
		if err != nil || !ok {
			log.Println("error while updating place and saving new background image: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update place: " + err.Error(),
			})
			return
		}

		log.Println("place updated and new background is saved..", placeID)
		c.JSON(200, "place updated and new background image upload complete")
	}
}
