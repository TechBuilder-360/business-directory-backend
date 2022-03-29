package repository

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/google/uuid"
	"time"
)

func (r *DefaultRepo) AddActivity(data *models.Activity) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	activity := &models.Activity{
		ID:        uuid.New().String(),
		Message:   data.Message,
		By:        data.By,
		For:       data.For,
		CreatedAt: time.Now().Local(),
	}
	_, err := r.Activities.InsertOne(ctx, &activity)
	if err != nil {
		return err
	}

	return nil
}
