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

	log.Println("Creating contraints...")
	err = db.Migrator().CreateConstraint(&models.Place{}, "Birds")
	if err != nil {
		log.Fatal("failed to create contraint on birds table : ", err)
	}
	err = db.Migrator().CreateConstraint(&models.Place{}, "fk_places_birds")
	if err != nil {
		log.Fatal("failed to create contraint table : ", err)
	}

	// check database foreign key for user & credit_cards exists or not
	ok := db.Migrator().HasConstraint(&models.Place{}, "Birds")
	if !ok {
		log.Fatal("nah")
	}
	ok = db.Migrator().HasConstraint(&models.Place{}, "fk_places_birds")
	if !ok {
		log.Fatal("nah 1")
	}

	log.Println("Migrating tables complete.")
}

//   // create database foreign key for user & credit_cards
//   db.Migrator().CreateConstraint(&models.Bird{}, "CreditCards")
//   db.Migrator().CreateConstraint(&models.Bird{}, "fk_users_credit_cards")
