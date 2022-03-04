package models

import "gorm.io/gorm"

type LandscapePlace struct {
	gorm.Model
	Name       string
	ShortName  string
	BGImg      string
	Landscapes []Landscape
}

type Landscape struct {
	gorm.Model
	LandscapePlaceID uint
	Quality          int
}

// TODO: migrate to WildlifePlace
type Place struct {
	gorm.Model
	Name      string
	ShortName string
	BGImg     string
	Birds     []Bird
}

// TODO: migrate to Wildlife
type Bird struct {
	gorm.Model
	Name    string
	PlaceID uint
	Quality int
}
