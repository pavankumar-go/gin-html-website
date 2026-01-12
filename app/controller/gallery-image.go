package controller

import (
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-html-website/database"
	"github.com/gin-html-website/models"
	"github.com/h2non/bimg"
)

func AddImage(name string, placeID uint, file *multipart.FileHeader, quality int) (*models.Bird, error) {
	image := &models.Bird{
		Name:    name,
		PlaceID: placeID,
		Quality: quality,
	}

	db := database.GetDBConnection()
	err := db.Create(image).Error
	if err != nil {
		log.Println("error adding bird..", err)
		return image, err
	}

	fileExtension := filepath.Ext(file.Filename)
	birdID := image.ID

	dstImgPath := fmt.Sprintf("static/assets/images/places/wildlife/%d/%d%s", placeID, birdID, fileExtension)
	dstImgThumbnailPath := fmt.Sprintf("static/assets/images/places/wildlife/%d/%d_thumb%s", placeID, birdID, fileExtension)

	dst, err := os.Create(dstImgPath)
	if err != nil {
		return image, err
	}
	defer dst.Close()

	fImg, err := file.Open()
	if err != nil {
		return image, err
	}
	defer fImg.Close()

	deimage, err := jpeg.Decode(fImg)
	if err != nil {
		log.Println("failed to decode image: ", err)
		return image, err
	}

	imgOpts := &jpeg.Options{
		Quality: quality,
	}

	err = jpeg.Encode(dst, deimage, imgOpts)
	if err != nil {
		log.Println("error saving full res image : ", err)
		return image, err
	}

	thumbImg, err := file.Open()
	if err != nil {
		return image, err
	}
	defer thumbImg.Close()

	imageBuf, err := io.ReadAll(thumbImg)
	if err != nil {
		log.Println("error reading image: ", err)
		return image, err
	}

	thumbnail, err := bimg.NewImage(imageBuf).Thumbnail(800)
	if err != nil {
		log.Println("error creating image thumbnail: ", err)
		return image, err
	}

	err = os.WriteFile(dstImgThumbnailPath, thumbnail, 0664)
	if err != nil {
		log.Println("error saving image thumbnail: ", err)
		return image, err
	}

	return image, nil
}

func RemoveImage(bID uint) (bool, error) {
	db := database.GetDBConnection()
	err := db.Unscoped().Delete(&models.Bird{}, bID).Error
	if err != nil {
		log.Println("error deleting image..", err)
		return false, err
	}

	log.Println("image deleted")
	return true, nil
}

func GetImages(placeId int) (*[]models.Bird, error) {
	db := database.GetDBConnection()
	var images []models.Bird
	err := db.Order("updated_at DESC").Where("place_id = ?", placeId).Find(&images).Having("place_id = ", placeId).Error // show latest first
	if err != nil {
		log.Println("error fetching images..", err)
		return nil, err
	}
	return &images, nil
}

func GetLatestUploadDate(placeID uint) time.Time {
	db := database.GetDBConnection()
	var image models.Bird
	err := db.Last(&image, "place_id = ?", placeID).Error // get latest upload date for a place
	if err != nil {
		log.Println("error getting latest images by upload date..", err)
	}

	return image.UpdatedAt
}
