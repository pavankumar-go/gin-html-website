package database

import (
	"log"

	"github.com/gin-html-website/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("Migrating tables...")
	err := db.AutoMigrate(&models.Bird{})
	if err != nil {
		log.Fatal("failed to migrate birds table : ", err)
	}

	err = db.AutoMigrate(&models.Place{})
	if err != nil {
		log.Fatal("failed to migrate birds table : ", err)
	}

	err = db.AutoMigrate(&models.Landscape{})
	if err != nil {
		log.Fatal("failed to migrate landscape table : ", err)
	}

	err = db.AutoMigrate(&models.LandscapePlace{})
	if err != nil {
		log.Fatal("failed to migrate landscape places table : ", err)
	}

	log.Println("Creating contraints...")
	err = db.Migrator().CreateConstraint(&models.Place{}, "Birds")
	if err != nil {
		log.Fatal("failed to create contraint on birds table : ", err)
	}
	err = db.Migrator().CreateConstraint(&models.Place{}, "fk_places_birds")
	if err != nil {
		log.Fatal("failed to create contraint table : ", err)
	}

	ok := db.Migrator().HasConstraint(&models.Place{}, "Birds")
	if !ok {
		log.Fatal("constraint check failed on Birds table")
	}

	ok = db.Migrator().HasConstraint(&models.Place{}, "fk_places_birds")
	if !ok {
		log.Fatal("constraint check failed on fk_places_birds")
	}

	// landscape
	err = db.Migrator().CreateConstraint(&models.LandscapePlace{}, "Landscapes")
	if err != nil {
		log.Fatal("failed to create contraint on Landscapes table : ", err)
	}

	//LandscapePlace ->	Landscape = fk_landscape_places_landscapes
	err = db.Migrator().CreateConstraint(&models.LandscapePlace{}, "fk_landscape_places_landscapes")
	if err != nil {
		log.Fatal("failed to create contraint table : ", err)
	}

	ok = db.Migrator().HasConstraint(&models.LandscapePlace{}, "Landscapes")
	if !ok {
		log.Fatal("constraint check failed on Landscapes table")
	}

	ok = db.Migrator().HasConstraint(&models.LandscapePlace{}, "fk_landscape_places_landscapes")
	if !ok {
		log.Fatal("constraint check failed on fk_landscape_places_landscapes")
	}

	log.Println("Migrating tables complete.")
}
