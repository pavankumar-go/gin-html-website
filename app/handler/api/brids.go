package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/controller"
)

func AddBird() gin.HandlerFunc {
	return func(c *gin.Context) {

		birdName, ok := c.GetPostForm("name")
		if !ok {
			log.Println("bird name missing")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing Bird Name",
			})
			return
		}

		placeIDStr, ok := c.GetPostForm("place_id")
		if !ok {
			log.Println("place id missing")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Missing Place ID",
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

		log.Println("saving bird image: ", birdName, "with place ID", placeID)
		bird, cErr := controller.AddBird(birdName, uint(placeID), fileHeader)
		if cErr != nil {
			log.Println("error occured while adding bird: cleanup in progress: ", err)
			deletedBird, err := controller.RemoveBird(bird.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "request was unsuccessful, cleanup failed too",
				})
				return
			}
			log.Println("cleanup completed..", deletedBird)

			log.Println("bird image saving had failed: ", cErr)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": cErr.Error(),
			})
			return
		}

		log.Println("image saved..", bird)
		c.JSON(200, "image upload complete")
	}
}
