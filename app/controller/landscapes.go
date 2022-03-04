package controller

import (
	"fmt"
	"image/jpeg"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-html-website/database"
	"github.com/gin-html-website/models"
)

func AddLandscape(placeID uint, file *multipart.FileHeader, quality int) (*models.Landscape, error) {
	landscape := &models.Landscape{
		LandscapePlaceID: placeID,
		Quality:          quality,
	}

	db := database.GetDBConnection()
	err := db.Create(landscape).Error
	if err != nil {
		log.Println("error adding landscape..", err)
		return landscape, err
	}

	fileExtension := filepath.Ext(file.Filename)
	landscapeID := landscape.ID

	dst, err := os.Create(fmt.Sprintf("static/assets/images/places/landscapes/%d/%d%s", placeID, landscapeID, fileExtension))
	if err != nil {
		return landscape, err
	}
	defer dst.Close()

	f, err := file.Open()
	if err != nil {
		return landscape, err
	}

	image, err := jpeg.Decode(f)
	if err != nil {
		log.Println("failed to decode image: ", err)
		return landscape, err
	}

	imgOpts := &jpeg.Options{
		Quality: quality,
	}

	err = jpeg.Encode(dst, image, imgOpts)
	if err != nil {
		log.Println("error saving image : ", err)
		return landscape, err
	}

	return landscape, nil
}

func GetLandscapes(placeId int) (*[]models.Landscape, error) {
	db := database.GetDBConnection()
	var landscapes []models.Landscape
	err := db.Order("updated_at DESC").Where("landscape_place_id = ?", placeId).Find(&landscapes).Having("landscape_place_id = ", placeId).Error // show latest first
	if err != nil {
		log.Println("error listing birds..", err)
		return nil, err
	}
	return &landscapes, nil
}

// TODO: FIXED
// 1. need to consider the re-ordering of image names `1.jpg 3.jpg ...`
// 2. OR find better way to scroll images than from index - current scroll 0,1,3 -> should work if 1,4,6,7,8,11

func RemoveLandscape(landscapeID uint) (bool, error) {
	db := database.GetDBConnection()
	err := db.Unscoped().Delete(&models.Landscape{}, landscapeID).Error
	if err != nil {
		log.Println("error deleting landscape..", err)
		return false, err
	}

	log.Printf("landscape %d deleted\n", landscapeID)
	return true, nil
}

func GetLatestUploadDateForLandscape(placeID uint) time.Time {
	db := database.GetDBConnection()
	var landscape models.Landscape
	err := db.Last(&landscape, "landscape_place_id = ?", placeID).Error // get latest upload date for a place
	if err != nil {
		log.Println("error getting lastest upload date..", err)
	}

	return landscape.UpdatedAt
}
