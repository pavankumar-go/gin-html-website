package api

import (
	"log"
	"net/http"

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
