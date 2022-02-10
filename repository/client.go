package repository

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (r *DefaultRepo) GetClientByID(cid string) ( *models.Client , error) {
	ctx, cancel:= context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	result:=models.Client{}
	filter := bson.D{{"ClientID", cid}}

	err := r.Client.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return &result, err
	}

	return &result, nil
}
