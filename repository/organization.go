package repository

import (
	"context"

	"time"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *DefaultRepo) CreateBranch(br *dto.CreateBranch) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	org := models.Branch{
		ID:               uuid.New().String(),
		OrganisationID: br.OrganisationID,
		BranchName: br.BranchName,
		Contact: br.Contact,   
		IsHQ: br.IsHQ,
		Address: br.Address,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}
	
	result, err := r.Branch.InsertOne(ctx, &org)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}

func (r *DefaultRepo) GetOrganisation() (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()


	cursor, err := r.Organisation.Find(ctx,bson.M{})
	

	return cursor, err
}

func (r *DefaultRepo) GetBranch(organisation string) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cursor, err := r.Branch.Find(ctx,bson.M{"organisation_id":organisation})
	

	return cursor, err
}

func (r *DefaultRepo) AlreadyOrganisation(br *dto.CreateOrganisation) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()


	err := r.Organisation.FindOne(ctx,bson.M{"organisation_name":br.OrganisationName}).Decode(&br)
	if err != nil {
		return false
	}

	return true
}

func (r *DefaultRepo) AlreadyBranch(br *dto.CreateBranch) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()


	err := r.Branch.FindOne(ctx,bson.M{"branch_name":br.BranchName}).Decode(&br)
	if err != nil {
		return false
	}

	return true
}

func (r *DefaultRepo) DeactivateOrganisation(id string,value bool) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter:=bson.M{"_id": id}
	update := bson.M{"$set":bson.M{"active": value}}
        result, err := r.Organisation.UpdateOne(ctx,filter , update)
	if err != nil {
		return nil, err
	}
	return result, nil
}
