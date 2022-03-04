package types

import "github.com/gin-html-website/models"

// blogs.html
// type Blogs struct {
// 	Blog []Blog
// }

// type Blog struct {
// 	Title   string
// 	Date    string
// 	Content string
// }

// places.html
type LandscapePlaces struct {
	LandscapePlace []models.LandscapePlace
}

// birds/place.html
type Landscapes struct {
	Landscape []models.Landscape
}

// places.html
type WildlifePlaces struct {
	Place []models.Place
}

// birds/place.html
type Wildlife struct {
	Bird []models.Bird
}
