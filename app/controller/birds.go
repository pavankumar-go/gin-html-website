package controller

import (
	"log"

	"github.com/gin-html-website/database"
	"github.com/gin-html-website/models"
)

func AddBird(name string) (*models.Bird, error) {
	data := &models.Bird{
		Name:    name,
		PlaceID: 100,
	}

	db := database.GetDBConnection()
	err := db.Create(data).Error
	if err != nil {
		log.Println("error adding bird..", err)
	}
	return nil, nil
}

func GetBirds() (*[]models.Bird, error) {
	db := database.GetDBConnection()
	var birds []models.Bird
	tx := db.Find(&birds)
	// rowsAffected := tx.RowsAffected
	err := tx.Error
	if err != nil {
		log.Println("error listing places..", err)
		return nil, err
	}

	for _, v := range birds {
		log.Println(v.ID, "---", v.Name)
	}

	return &birds, nil
}
