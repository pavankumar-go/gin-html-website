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
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")
		c.HTML(200, "gallery.html", nil)
	}
}

func WildlifePlaces() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")
		places, err := controller.GetWildlifePlaces()
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		var allPlaces types.WildlifePlaces

		for _, place := range *places {
			place.UpdatedAt = controller.GetLatestUploadDate(place.ID)
			allPlaces.Place = append(allPlaces.Place, place)
		}

		tmpl := template.Must(template.ParseFiles("templates/main/wildlife_places.html"))
		err = tmpl.Execute(c.Writer, allPlaces)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
	}
}

// TODO: update
func LandscapePlaces() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")
		places, err := controller.GetLandscapePlaces()
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		var allPlaces types.LandscapePlaces

		for _, place := range *places {
			place.UpdatedAt = controller.GetLatestUploadDateForLandscape(place.ID)
			allPlaces.LandscapePlace = append(allPlaces.LandscapePlace, place)
		}

		tmpl := template.Must(template.ParseFiles("templates/main/landscape_places.html"))
		err = tmpl.Execute(c.Writer, allPlaces)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
	}
}

// wildlife in places handlers...
func W_Bangalore() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")

		birds, err := controller.GetBirds(1) // 1: bangalore
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
		tmpl := template.Must(template.ParseFiles("templates/main/places/wildlife/common.html"))
		err = tmpl.Execute(c.Writer, allBirds)
		if err != nil {
			log.Println("failed to render birds: ", err)
			c.AbortWithStatusJSON(500, "failed to render birds")
			return
		}
	}
}

func W_Mandya() gin.HandlerFunc {
	return func(c *gin.Context) {
		birds, err := controller.GetBirds(2) // 2: mandya
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
		tmpl := template.Must(template.ParseFiles("templates/main/places/wildlife/common.html"))
		err = tmpl.Execute(c.Writer, allBirds)
		if err != nil {
			log.Println("failed to render birds: ", err)
			c.AbortWithStatusJSON(500, "failed to render birds")
			return
		}
	}
}

// wildlife in places handlers...
func L_Bangalore() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")

		landscapes, err := controller.GetLandscapes(1) // 1: bangalore
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		allLandscapes := types.Landscapes{
			Landscape: *landscapes,
		}

		// NOTE: this should pre-exist corresponds to models.ShortName
		// Even API should comply to this shortName /places/landscape/<shortName>
		tmpl := template.Must(template.ParseFiles("templates/main/places/landscapes/common.html"))
		err = tmpl.Execute(c.Writer, allLandscapes)
		if err != nil {
			log.Println("failed to render landscapes: ", err)
			c.AbortWithStatusJSON(500, "failed to render landscapes")
			return
		}
	}
}

func L_Mandya() gin.HandlerFunc {
	return func(c *gin.Context) {
		landscapes, err := controller.GetLandscapes(2) // 2: mandya
		if err != nil {
			log.Println("failed to get places: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}

		allLandscapes := types.Landscapes{
			Landscape: *landscapes,
		}

		// NOTE: this should pre-exist corresponds to models.ShortName
		// Even API should comply to this shortName /places/landscapes/<shortName>
		tmpl := template.Must(template.ParseFiles("templates/main/places/landscapes/common.html"))
		err = tmpl.Execute(c.Writer, allLandscapes)
		if err != nil {
			log.Println("failed to render landscapes: ", err)
			c.AbortWithStatusJSON(500, "failed to render landscapes")
			return
		}
	}
}
