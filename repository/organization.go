package repository

import (
	"context"
	"time"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/google/uuid"
)

func (r *DefaultRepo) CreateOrganisation(Organs *dto.CreateOrganisation) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	org := models.Organisation{
		ID:               uuid.New().String(),
		OrganisationName: Organs.OrganisationName,
		OrganisationSize: Organs.OrganisationSize,
		FoundingDate:     Organs.FoundingDate,
		Description:      Organs.Description,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}

	result, err := r.Organisation.InsertOne(ctx, &org)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}
