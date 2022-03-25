package repository

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"time"
)

func (r *DefaultRepo) AddActivity(data *models.Activity) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := r.Activities.InsertOne(ctx, &data)
	if err != nil {
		return err
	}

	return nil
}
