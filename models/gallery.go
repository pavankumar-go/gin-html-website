package models

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	Places []Place
}
