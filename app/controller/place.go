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
	"gorm.io/gorm"
)

func AddPlace(name, shortName string, file *multipart.FileHeader) (*models.Place, error) {
	db := database.GetDBConnection()
	place := &models.Place{
		Name:      name,
		ShortName: shortName,
	}

	err := db.Create(place).Error
	if err != nil {
		log.Println("error adding place..", err)
		return place, err
	}
	fileExtension := filepath.Ext(file.Filename)
	placeID := place.ID

	err = os.MkdirAll(fmt.Sprintf("static/assets/images/places/%d/bg", placeID), 0777)
	if err != nil {
		log.Println("error creating directory..", err)
		return place, err
	}

	dst, err := os.Create(fmt.Sprintf("static/assets/images/places/%d/bg/%d-place-bg%s", placeID, placeID, fileExtension))
	if err != nil {
		log.Println("error saving background image..", err)
		return place, err
	}

	defer dst.Close()

	image, err := file.Open()
	if err != nil {
		return place, err
	}

	_, err = io.Copy(dst, image)
	if err != nil {
		log.Println("error saving attachment place image background: ", err)
		return place, err
	}

	placeBGImg := &models.Place{
		BGImg: fmt.Sprintf("/%d/bg/%d-place-bg%s", placeID, placeID, fileExtension),
	}

	err = db.Model(models.Place{}).Where("id = ?", placeID).Updates(placeBGImg).Error
	if err != nil {
		log.Println("error updating place for image background path: ", err)
		return place, err
	}

	return place, nil
}

func GetPlaces() (*[]models.Place, error) {
	db := database.GetDBConnection()
	var places []models.Place
	err := db.Find(&places).Error
	if err != nil {
		log.Println("error getting places..", err)
		return nil, err
	}

	return &places, nil
}

func RemovePlace(placeID uint) (bool, error) {
	db := database.GetDBConnection()
	err := db.Unscoped().Delete(&models.Place{}, placeID).Error
	if err != nil {
		log.Println("error deleting place..", err)
		return false, err
	}

	pathToDel := fmt.Sprintf("static/assets/images/places/%d", placeID)
	err = os.RemoveAll(pathToDel)
	if err != nil {
		log.Printf("error deleting place directory %s: %v", pathToDel, err)
		return false, err
	}

	log.Println("permanently deleted place: ", placeID)
	return true, nil
}

func UpdatePlace(placeID uint, file *multipart.FileHeader) (bool, error) {
	db := database.GetDBConnection()

	fileExtension := filepath.Ext(file.Filename)

	place := &models.Place{
		Model: gorm.Model{
			ID: placeID,
		},
	}

	err := db.First(&place).Error
	log.Println("DB error:", err)
	if err != nil {
		log.Println("error updating place for image background path: ", err)
		return false, err
	}

	dst, err := os.Create(fmt.Sprintf("static/assets/images/places/%d/bg/%d-place-bg%s", placeID, placeID, fileExtension))
	if err != nil {
		log.Println("error saving new background image..", err)
		return false, err
	}

	defer dst.Close()

	image, err := file.Open()
	if err != nil {
		log.Println("failed to open image: ", err)
		return false, err
	}

	_, err = io.Copy(dst, image)
	if err != nil {
		log.Println("error saving attachment for new place image background: ", err)
		return false, err
	}

	return true, nil
}
