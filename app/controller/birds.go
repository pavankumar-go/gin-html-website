package controller

import (
	"fmt"
	"image/jpeg"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-html-website/database"
	"github.com/gin-html-website/models"
)

func AddBird(name string, placeID uint, file *multipart.FileHeader, quality int) (*models.Bird, error) {
	bird := &models.Bird{
		Name:    name,
		PlaceID: placeID,
		Quality: quality,
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

	f, err := file.Open()
	if err != nil {
		return bird, err
	}

	image, err := jpeg.Decode(f)
	if err != nil {
		log.Println("failed to decode image: ", err)
		return bird, err
	}

	imgOpts := &jpeg.Options{
		Quality: quality,
	}

	err = jpeg.Encode(dst, image, imgOpts)
	if err != nil {
		log.Println("error saving image : ", err)
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

// TODO:
// 1. need to consider the re-ordering of image names `1.jpg 3.jpg ...`
// 2. OR find better way to scroll images than from index - current scroll 0,1,3 -> should work if 1,4,6,7,8,11

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
