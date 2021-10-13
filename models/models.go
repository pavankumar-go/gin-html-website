package models

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Name      string
	ShortName string
	BGImg     string
	Birds     []Bird
}

type Bird struct {
	gorm.Model
	Name    string
	PlaceID uint
}
