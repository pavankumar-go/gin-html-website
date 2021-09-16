package handler

import "github.com/gin-gonic/gin"

func HomePage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "home_page.html", nil)
	}
}

func About() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Blogs() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Places() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
