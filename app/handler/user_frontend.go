package handler

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/controller"
	"github.com/gin-html-website/app/types"
	// "github.com/gin-html-website/app/types"
)

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")
		c.HTML(200, "home.html", nil)
	}
}

func About() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")
		c.HTML(200, "about.html", nil)
	}
}

func Blogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")
		tmpl := template.Must(template.ParseFiles("templates/main/blogs.html"))
		err := tmpl.Execute(c.Writer, "blogs")
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
	}
}

func Places() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")
		places, err := controller.GetPlaces()
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		var allPlaces types.Places

		for _, place := range *places {
			place.UpdatedAt = controller.GetLatestUploadDate(place.ID)
			allPlaces.Place = append(allPlaces.Place, place)
		}

		tmpl := template.Must(template.ParseFiles("templates/main/places.html"))
		err = tmpl.Execute(c.Writer, allPlaces)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
	}
}

// Birds in places handlers...

func Bangalore() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")

		birds, err := controller.GetBirds(1) // 1: bangalore
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		allBirds := types.Birds{
			Bird: *birds,
		}

		// NOTE: this should pre-exist corresponds to models.ShortName
		// Even API should comply to this shortName /places/<shortName>
		tmpl := template.Must(template.ParseFiles("templates/birds/common.html"))
		err = tmpl.Execute(c.Writer, allBirds)
		if err != nil {
			log.Println("failed to render birds: ", err)
			c.AbortWithStatusJSON(500, "failed to render birds")
			return
		}
	}
}

func Mandya() gin.HandlerFunc {
	return func(c *gin.Context) {
		birds, err := controller.GetBirds(2) // 2: mandya
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		allBirds := types.Birds{
			Bird: *birds,
		}

		// NOTE: this should pre-exist corresponds to models.ShortName
		// Even API should comply to this shortName /places/<shortName>
		tmpl := template.Must(template.ParseFiles("templates/birds/common.html"))
		err = tmpl.Execute(c.Writer, allBirds)
		if err != nil {
			log.Println("failed to render birds: ", err)
			c.AbortWithStatusJSON(500, "failed to render birds")
			return
		}
	}
}
