package migration

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Seed the database with some data
func Seed(db *gorm.DB) {
	var errs []error
	errs = append(errs, runCategorySeeder(db))
	errs = append(errs, runRolesSeeder(db))
	errs = append(errs, runCountrySeeder(db))

	for _, e := range errs {
		if e != nil {
			log.Errorf("migration apiError-> %v", e)
		}
	}
}

func runCategorySeeder(tx *gorm.DB) error {
	//categories := []model.Category{
	//	{
	//		Name: "Information Technology",
	//	},
	//	{
	//		Name: "Commerce",
	//	},
	//}
	//
	//if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&categories).Error; err != nil {
	//	return err
	//}

	return nil
}

func runRolesSeeder(tx *gorm.DB) error {
	//roles := []Role{
	//	{
	//		Name: "Owner",
	//	},
	//	{
	//		Name: "Organisation Admin",
	//	},
	//	{
	//		Name: "Branch Manager",
	//	},
	//}
	//
	//if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&roles).Error; err != nil {
	//	return err
	//}

	return nil
}

func runCountrySeeder(tx *gorm.DB) error {
	//country := []model2.Country{
	//	{
	//		Name:        "Nigeria",
	//		Code:        "NG",
	//		CallingCode: "234",
	//		Active:      true,
	//	},
	//}
	//
	//if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&country).Error; err != nil {
	//	return err
	//}

	return nil
}
