package seeder

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Seed the database with some data
func Seed(db *gorm.DB) {
	var errs []error
	err := runCategorySeeder(db)
	errs = append(errs, err)
	err = runRolesSeeder(db)
	errs = append(errs, err)

	for _, e := range errs {
		if e != nil {
			log.Errorf("seeder error-> %v", e)
		}
	}
}

func runCategorySeeder(tx *gorm.DB) error {
	categories := []model.Category{
		{
			Name: "Information Technology",
		},
		{
			Name: "Commerce",
		},
	}

	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories).Error; err != nil {
		return err
	}

	return nil
}

func runRolesSeeder(tx *gorm.DB) error {
	roles := []model.Role{
		{
			Name: "Owner",
		},
		{
			Name: "Organisation Admin",
		},
		{
			Name: "Branch Manager",
		},
	}

	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&roles).Error; err != nil {
		return err
	}

	return nil
}
