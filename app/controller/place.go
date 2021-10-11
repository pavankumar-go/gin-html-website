package controller

import (
	"log"

	"github.com/gin-html-website/database"
	"github.com/gin-html-website/models"
)

func AddPlace(name string) (*models.Place, error) {
	data := &models.Place{
		Name: name,
	}

	var place models.Place
	var count int64

	db := database.GetDBConnection()
	err := db.Where(data).First(&place).Count(&count).Error
	if err != nil {
		if err.Error() == "record not found" {
			if count == 0 {
				err = db.Model(&place).Create(data).Error
				if err != nil {
					return nil, err
				}
			}
			return nil, nil
		}
		return nil, err
	}
	return &place, nil
}

func GetPlaces() (*[]models.Place, error) {
	db := database.GetDBConnection()
	var places []models.Place
	tx := db.Find(&places)
	// rowsAffected := tx.RowsAffected
	err := tx.Error
	if err != nil {
		log.Println("error listing places..", err)
		return nil, err
	}

	for _, v := range places {
		log.Println(v.ID, "---", v.Name)
	}

	return &places, nil
}
