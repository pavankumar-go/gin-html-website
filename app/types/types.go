package types

type BirdsDTO struct {
	ID   string
	Name string
}

type PlaceDTO struct {
	ID   string
	Name string
}

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
	Place []Place
}

type Place struct {
	ID             string
	Name           string
	LastUpdateDate string
	Handler        string
}

// birds.html
type Birds struct {
	Bird []Bird
}

type Bird struct {
	ID   string
	Name string
}
