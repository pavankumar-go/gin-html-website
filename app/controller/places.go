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
	"github.com/h2non/bimg"
	"gorm.io/gorm"
)

func AddWildlifePlace(name, shortName string, file *multipart.FileHeader) (*models.Place, error) {
	db := database.GetDBConnection()
	place := &models.Place{
		Name:      name,
		ShortName: shortName,
	}

	err := db.Create(place).Update("place_id", gorm.Expr("id")).Error
	if err != nil {
		log.Println("error adding place..", err)
		return place, err
	}

	fileExtension := ".jpg"
	placeID := place.ID

	err = os.MkdirAll(fmt.Sprintf("static/assets/images/places/wildlife/%d/bg", placeID), 0777)
	if err != nil {
		log.Println("error creating directory..", err)
		return place, err
	}

	image, err := file.Open()
	if err != nil {
		return place, err
	}

	imageBuf, err := io.ReadAll(image)
	if err != nil {
		log.Println("error reading place image: ", err)
		return place, err
	}

	thumbnail, err := bimg.NewImage(imageBuf).Thumbnail(1000)
	if err != nil {
		log.Println("error creating place image thumbnail: ", err)
		return place, err
	}

	dstFilePath := fmt.Sprintf("static/assets/images/places/wildlife/%d/bg/%d-place-bg%s", placeID, placeID, fileExtension)
	err = os.WriteFile(dstFilePath, thumbnail, 0664)
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

func AddLandscapePlace(name, shortName string, file *multipart.FileHeader) (*models.LandscapePlace, error) {
	db := database.GetDBConnection()
	place := &models.LandscapePlace{
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

	err = os.MkdirAll(fmt.Sprintf("static/assets/images/places/landscapes/%d/bg", placeID), 0777)
	if err != nil {
		log.Println("error creating directory..", err)
		return place, err
	}

	dst, err := os.Create(fmt.Sprintf("static/assets/images/places/landscapes/%d/bg/%d-place-bg%s", placeID, placeID, fileExtension))
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

	landscapePlaceBGImg := &models.LandscapePlace{
		BGImg: fmt.Sprintf("/%d/bg/%d-place-bg%s", placeID, placeID, fileExtension),
	}

	err = db.Model(models.LandscapePlace{}).Where("id = ?", placeID).Updates(landscapePlaceBGImg).Error
	if err != nil {
		log.Println("error updating place for image background path: ", err)
		return place, err
	}

	return place, nil
}

func GetWildlifePlaces() (*[]models.Place, error) {
	db := database.GetDBConnection()
	var places []models.Place
	err := db.Find(&places).Error
	if err != nil {
		log.Println("error getting places..", err)
		return nil, err
	}

	return &places, nil
}

func GetPlacesByLatestUploads() (*[]models.Place, error) {
	db := database.GetDBConnection()
	var places []models.Place
	err := db.Raw("select places.*, max(birds.updated_at) from places join birds on birds.place_id = places.place_id group by birds.place_id order by birds.updated_at desc;").Scan(&places).Error
	if err != nil {
		log.Println("error getting places..", err)
		return nil, err
	}

	return &places, nil
}

func RemoveWildlifePlace(placeID uint) (bool, error) {
	db := database.GetDBConnection()
	err := db.Unscoped().Delete(&models.Place{}, placeID).Error
	if err != nil {
		log.Println("error deleting place..", err)
		return false, err
	}

	pathToDel := fmt.Sprintf("static/assets/images/places/wildlife/%d", placeID)
	err = os.RemoveAll(pathToDel)
	if err != nil {
		log.Printf("error deleting place directory %s: %v", pathToDel, err)
		return false, err
	}

	log.Println("permanently deleted place: ", placeID)
	return true, nil
}
