package database

import (
	"log"

	"github.com/gin-html-website/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("Seeding place table...")
	places := []string{"Valley School", "Jogi Kere", "Turahalli"}

	// var place models.Place
	// tx := db.First(&place, models.Place{})
	// if tx.RowsAffected == 0 {
	for id, place := range places {
		addPlace(place, id)
	}
	log.Println("Seeding place table complete...")
	// }
}

func addPlace(name string, id int) (*models.Place, error) {
	log.Println(name, "---", id)
	var place models.Place
	db := GetDBConnection()
	data := &models.Place{
		Name: name,
	}

	err := db.Create(data).Error
	if err != nil {
		log.Fatalln(err)
	}
	return &place, nil
}
