package controller

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-html-website/database"
	"github.com/gin-html-website/models"
)

func AddBird(name string, placeID uint, file *multipart.FileHeader) (*models.Bird, error) {
	bird := &models.Bird{
		Name:    name,
		PlaceID: placeID,
	}

	db := database.GetDBConnection()
	err := db.Create(bird).Error
	if err != nil {
		log.Println("error adding bird..", err)
		return bird, err
	}

	fileExtension := filepath.Ext(file.Filename)
	birdID := bird.ID

	dst, err := os.Create(fmt.Sprintf("static/assets/images/places/%d/%d%s", placeID, birdID, fileExtension))
	if err != nil {
		return bird, err
	}
	defer dst.Close()

	image, err := file.Open()
	if err != nil {
		return bird, err
	}

	_, err = io.Copy(dst, image)
	if err != nil {
		log.Println("error saving attachment image: ", err)
		return bird, err
	}

	return bird, nil
}

func GetBirds() (*[]models.Bird, error) {
	db := database.GetDBConnection()
	var birds []models.Bird
	err := db.Find(&birds).Error
	if err != nil {
		log.Println("error listing birds..", err)
		return nil, err
	}

	return &birds, nil
}

func RemoveBird(bID uint) (bool, error) {
	db := database.GetDBConnection()
	err := db.Unscoped().Delete(&models.Bird{}, bID).Error
	if err != nil {
		log.Println("error deleting bird..", err)
		return false, err
	}

	log.Println("bird deleted")
	return true, nil
}
