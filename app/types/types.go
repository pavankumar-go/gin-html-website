package types

import "github.com/gin-html-website/models"

// blogs.html
type Blogs struct {
	Blog []Blog
}

type Blog struct {
	Title   string
	Date    string
	Content string
}

// places.html
type Places struct {
	Place []models.Place
}

// birds/place.html
type Birds struct {
	Bird []models.Bird
}
