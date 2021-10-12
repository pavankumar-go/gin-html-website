package handler

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-html-website/app/types"
)

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		c.HTML(200, "home.html", nil)
	}
}

func About() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		c.HTML(200, "about.html", nil)
	}
}

func Blogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		blogs := types.Blogs{
			Blog: []types.Blog{
				{
					Date:  "21 Oct 2021",
					Title: "Adventures of valley school",
					Content: `
				This is a great space to write long text about your company and your services. You can use this
                space to go into a little more detail about your company. Talk about your team and what services you
                provide. Tell your visitors the story of how you came up with the idea for your business and what
                makes you different from your competitors. Make your company stand out and show your visitors who
                you are. Tip: Add your own image by double clicking the image and clicking Change Image`,
				},
				{
					Date:  "21 Nov 2021",
					Title: "Trip to KP",
					Content: `
				This is a great space to write long text about your company and your services. You can use this
                space to go into a little more detail about your company. Talk about your team and what services you
                provide. Tell your visitors the story of how you came up with the idea for your business and what
                makes you different from your competitors. Make your company stand out and show your visitors who
                you are. Tip: Add your own image by double clicking the image and clicking Change Image`,
				},
				{
					Date:  "21 Dec 2021",
					Title: "Photo Blinds",
					Content: `
				This is a great space to write long text about your company and your services. You can use this
                space to go into a little more detail about your company. Talk about your team and what services you
                provide. Tell your visitors the story of how you came up with the idea for your business and what
                makes you different from your competitors. Make your company stand out and show your visitors who
                you are. Tip: Add your own image by double clicking the image and clicking Change Image`,
				},
			},
		}

		tmpl := template.Must(template.ParseFiles("templates/main/blogs.html"))
		err := tmpl.Execute(c.Writer, blogs)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
		// c.HTML(200, "blogs.html", nil)
		// c.Writer
	}
}

func Places() gin.HandlerFunc {
	return func(c *gin.Context) {
		places := types.Places{
			Place: []types.Place{
				{
					ID:             "1",
					LastUpdateDate: "21 Oct 2021",
					Name:           "Bangalore",
					Handler:        "/places/bangalore",
				},
				{
					ID:             "2",
					LastUpdateDate: "21 Nov 2021",
					Name:           "Valley School",
					Handler:        "/places/bangalore",
				},
				{
					ID:             "3",
					LastUpdateDate: "21 Dev 2021",
					Name:           "OMH",
					Handler:        "/places/bangalore",
				},
			},
		}
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")

		tmpl := template.Must(template.ParseFiles("templates/main/places.html"))
		err := tmpl.Execute(c.Writer, places)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
		// c.Header("max-age", "0")
		// c.Header("Cache-Control", "no-cache")

		// c.HTML(200, "places.html", nil)
	}
}

func Bangalore() gin.HandlerFunc {
	return func(c *gin.Context) {
		birds := types.Birds{
			Bird: []types.Bird{
				{
					ID:   "1",
					Name: "bird-name",
				},
				{
					ID:   "2",
					Name: "bird-name",
				},
				{
					ID:   "3",
					Name: "bird-name",
				},
				{
					ID:   "4",
					Name: "bird-name",
				},
				{
					ID:   "5",
					Name: "bird-name",
				},
				{
					ID:   "6",
					Name: "bird-name",
				},
				{
					ID:   "7",
					Name: "bird-name",
				},
				{
					ID:   "8",
					Name: "bird-name",
				},
				{
					ID:   "9",
					Name: "bird-name",
				},
				{
					ID:   "4",
					Name: "bird-name",
				},
				{
					ID:   "5",
					Name: "bird-name",
				},
				{
					ID:   "6",
					Name: "bird-name",
				},
			},
		}
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")

		tmpl := template.Must(template.ParseFiles("templates/places/bangalore.html"))
		err := tmpl.Execute(c.Writer, birds)
		if err != nil {
			log.Println("failed to render: ", err)
			c.AbortWithStatusJSON(500, "failed to render")
			return
		}
	}
}

func Gaganachukki() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		c.HTML(200, "gaganachukki.html", nil)
	}
}

func Ganeshgudi() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		c.HTML(200, "ganesh_gudi.html", nil)
	}
}

func ValleySchool() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		c.HTML(200, "valley_school.html", nil)
	}
}
