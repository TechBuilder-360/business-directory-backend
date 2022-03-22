package repository

import (
	"context"
	"math"
	"strconv"

	"time"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)





func (r *DefaultRepo) CreateOrganisation(Organs *dto.CreateOrgReq) (*dto.CreateOrgResponse, error) {
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
		Active: true,
	}
	result, err := r.Organisation.InsertOne(ctx, &org)
	if err != nil {
		return nil, err
	}

	br:= models.Branch{
		ID:               uuid.New().String(),
		OrganisationID: result.InsertedID.(string),
		BranchName: Organs.OrganisationName,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
		IsHQ: true,

	}
	
	_, err = r.Branch.InsertOne(ctx, &br)

	if err != nil {
		return nil, err
	}

	return &dto.CreateOrgResponse{OrganisationID: result.InsertedID.(string)}, nil
}

func (r *DefaultRepo) CreateBranch(br *dto.CreateBranch) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	org := models.Branch{
		ID:               uuid.New().String(),
		OrganisationID: br.OrganisationID,
		BranchName: br.BranchName,
		Contact: br.Contact,   
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

func (r *DefaultRepo) GetOrganisation(page string) (*dto.DataView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var br []bson.M

	findOptions:=options.Find()
	pageP,_:=strconv.Atoi(page)
	total,_:=r.Organisation.CountDocuments(ctx,bson.M{})
	var perPage int64=10
	findOptions.SetLimit(int64(perPage))
	findOptions.SetSkip((int64(pageP)-1)* perPage)
	cursor, err := r.Organisation.Find(ctx,bson.M{},findOptions)
	err = cursor.All(ctx, &br)
	data:=&dto.DataView{
		Page:pageP,
		Perpage:perPage,
		Total:total,
		LastPage: math.Ceil(float64(total / perPage)),
                Data:br,

	}
	return data, err
}

func (r *DefaultRepo) GetBranch(organisation string,page string) (*dto.DataView,  error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var br []bson.M
	findOptions:=options.Find()
	pageP,_:=strconv.Atoi(page)
	total,_:=r.Branch.CountDocuments(ctx,bson.M{})
	var perPage int64=10
	findOptions.SetLimit(int64(perPage))
	findOptions.SetSkip((int64(pageP)-1)* perPage)
	cursor, err := r.Branch.Find(ctx,bson.M{"organisation_id":organisation},findOptions)
	_= cursor.All(ctx, &br)
	data:=&dto.DataView{
		Page:pageP,
		Perpage:perPage,
		Total:total,
		LastPage: math.Ceil(float64(total / perPage)),
                Data:br,

	}
	return data, err
}

func (r *DefaultRepo) AlreadyOrganisation(br *dto.CreateOrgReq) bool {
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

func (r *DefaultRepo) DeactivateOrganisation(br *dto.DeReactivateOrgReq) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter:=bson.M{"_id": br.OrganisationID}
	update := bson.M{"$set":bson.M{"active": br.Active}}
        result, err := r.Organisation.UpdateOne(ctx,filter , update)
	if err != nil {
		return nil, err
	}
	return result, nil
}
