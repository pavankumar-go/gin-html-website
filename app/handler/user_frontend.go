package handler

import (
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/controller"
	"github.com/gin-html-website/app/types"
)

func HomePage(appPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		slides := []string{}
		files, err := os.ReadDir(appPath + "/static/assets/images/slideshow")
		if err != nil {
			c.AbortWithStatusJSON(500, "failed to render")
		}

		for _, f := range files {
			slides = append(slides, "/static/assets/images/slideshow/"+f.Name())
		}

		tmpl := template.Must(template.ParseFiles("templates/main/home.html"))
		err = tmpl.Execute(c.Writer, slides)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
	}
}

func About() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "about.html", nil)
	}
}

// func Blogs() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// c.Header("max-age", "0")
// 		// c.Header("Cache-Control", "no-cache")
// 		tmpl := template.Must(template.ParseFiles("templates/main/blogs.html"))
// 		err := tmpl.Execute(c.Writer, "blogs")
// 		if err != nil {
// 			log.Println("failed to render: ", err)
// 			c.AbortWithStatusJSON(500, "failed to render")
// 			return
// 		}
// 	}
// }

func Gallery() gin.HandlerFunc {
	return func(c *gin.Context) {
		latestPlaces, err := controller.GetPlacesByLatestUploads()
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		var allPlaces types.WildlifePlaces
		for _, place := range *latestPlaces {
			place.UpdatedAt = controller.GetLatestUploadDate(place.ID)
			allPlaces.Place = append(allPlaces.Place, place)
		}

		tmpl := template.Must(template.ParseFiles("templates/main/collections.html"))
		err = tmpl.Execute(c.Writer, allPlaces)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
	}
}

// wildlife in places handlers...
func W_Places(placeID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		birds, err := controller.GetImages(placeID)
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		allBirds := types.Wildlife{
			Bird: *birds,
		}

		// NOTE: this should pre-exist corresponds to models.ShortName
		// Even API should comply to this shortName /places/<shortName>
		tmpl := template.Must(template.ParseFiles("templates/main/places/common.html"))
		err = tmpl.Execute(c.Writer, allBirds)
		if err != nil {
			log.Println("failed to render birds: ", err)
			c.AbortWithStatusJSON(500, "failed to render birds")
			return
		}
	}
}
