package handler

import "github.com/gin-gonic/gin"

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
		c.HTML(200, "blogs.html", nil)
	}
}

func Places() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		c.HTML(200, "places.html", nil)
	}
}

func Bangalore() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("max-age", "0")
		c.Header("Cache-Control", "no-cache")
		c.HTML(200, "bangalore.html", nil)
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
