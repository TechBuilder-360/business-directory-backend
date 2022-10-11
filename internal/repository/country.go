package repository

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
	"strings"
)

//go:generate mockgen -destination=../mocks/repository/country.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository CountryRepository
type CountryRepository interface {
	GetCountryByID(id string) (*model.Country, error)
	GetCountryByCode(code string) (*model.Country, error)
	WithTx(tx *gorm.DB) CountryRepository
}

type countryRepo struct {
	db *gorm.DB
}

func (c *countryRepo) GetCountryByID(id string) (*model.Country, error) {
	var country model.Country
	if err := c.db.Where("id = ?", id).First(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("country not found")
		}
		return nil, err
	}
	return &country, nil
}

func (c *countryRepo) GetCountryByCode(code string) (*model.Country, error) {
	var country model.Country
	if err := c.db.Where("code = ? and active = true", strings.ToUpper(code)).First(&country).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("country not found")
		}
		return nil, err
	}
	return &country, nil
}

func (c *countryRepo) WithTx(tx *gorm.DB) CountryRepository {
	return &countryRepo{
		db: tx,
	}
}

func NewCountryRepository() CountryRepository {
	return &countryRepo{
		db: database.ConnectDB(),
	}
}
