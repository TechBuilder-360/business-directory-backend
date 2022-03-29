package repository

import (
	"context"
	"fmt"
	// "encoding/json"
	// "fmt"
	"math"
	"strconv"

	"time"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *DefaultRepo) CreateOrganisation(Organs *dto.CreateOrgReq) (*dto.CreateOrgResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	
	returnID := ""
	session, err :=r.Cli.StartSession()
	if err != nil {
		return nil,err
	}
	defer session.EndSession(ctx)
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
		    return err
		}
		org := models.Organisation{
			ID:               uuid.New().String(),
			OrganisationName: Organs.OrganisationName,
			OrganisationSize: Organs.OrganisationSize,
			FoundingDate:     Organs.FoundingDate,
			Description:      Organs.Description,
			CreatedAt:        time.Now().Local(),
			UpdatedAt:        time.Now().Local(),
			Active:           true,
		}
		result, err := r.Organisation.InsertOne(sessionContext, &org)
		if err != nil {
			
			return err
		}

		br := models.Branch{
			ID:             uuid.New().String(),
			OrganisationID: result.InsertedID.(string),
			BranchName:     Organs.OrganisationName,
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			IsHQ:           true,
		}

		_, err = r.Branch.InsertOne(sessionContext, &br)
		if err != nil {
			
			session.AbortTransaction(sessionContext)
			return err
		}

		if err = session.CommitTransaction(sessionContext); err != nil {
			fmt.Print(result.InsertedID.(string))
		    return err
		}
		returnID = result.InsertedID.(string)
		return nil
	    })

	   if err != nil {
		   return nil,err
	   }
	   
		return &dto.CreateOrgResponse{OrganisationID: returnID}, nil

	}
	

func (r *DefaultRepo) CreateBranch(br *dto.CreateBranch) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	org := models.Branch{
		ID:             uuid.New().String(),
		OrganisationID: br.OrganisationID,
		BranchName:     br.BranchName,
		Contact:        br.Contact,
		Address:        br.Address,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	result, err := r.Branch.InsertOne(ctx, &org)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(string), nil
}

func (r *DefaultRepo) GetOrganisations(page string) (*dto.DataView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var br []bson.M

	findOptions := options.Find()
	pageP, _ := strconv.Atoi(page)
	total, _ := r.Organisation.CountDocuments(ctx, bson.M{})
	var perPage int64 = 10
	findOptions.SetProjection(bson.M{
		"active":0,
		"admins":0,
		"creator_id":0,

	})
	findOptions.SetLimit(perPage)
	findOptions.SetSkip((int64(pageP) - 1) * perPage)
	
	cursor, err := r.Organisation.Find(ctx, bson.M{}, findOptions)
	err = cursor.All(ctx,&br)

	data := &dto.DataView{
		Page:     pageP,
		Perpage:  perPage,
		Total:    total,
		LastPage: math.Ceil(float64(total / perPage)),
		Data:     br,
	}
	return data, err
}

func (r *DefaultRepo) GetSingleOrganisation(id string) (primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var br bson.M
	err := r.Organisation.FindOne(ctx, bson.M{"_id": id}).Decode(&br)

	return br, err
}

func (r *DefaultRepo) GetSingleBranch(id string) (primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var br bson.M
	err := r.Branch.FindOne(ctx, bson.M{"_id": id}).Decode(&br)
	return br, err
}
func (r *DefaultRepo) GetBranches(organisation string, page string) (*dto.DataView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var br []bson.M
	findOptions := options.Find()
	pageP, _ := strconv.Atoi(page)
	total, _ := r.Branch.CountDocuments(ctx, bson.M{})
	var perPage int64 = 10
	findOptions.SetProjection(bson.M{
		"IsHQ":0,
		"admins":0,
	})
	findOptions.SetLimit(int64(perPage))
	findOptions.SetSkip((int64(pageP) - 1) * perPage)
	cursor, err := r.Branch.Find(ctx, bson.M{"organisation_id": organisation}, findOptions)
	_ = cursor.All(ctx, &br)
	data := &dto.DataView{
		Page:     pageP,
		Perpage:  perPage,
		Total:    total,
		LastPage: math.Ceil(float64(total / perPage)),
		 Data:     br,
	}
	return data, err
}

func (r *DefaultRepo) OrganisationExist(br *dto.CreateOrgReq) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := r.Organisation.FindOne(ctx, bson.M{"organisation_name": br.OrganisationName}).Decode(&br)
	if err != nil {
		return false
	}

	return true
}

func (r *DefaultRepo) BranchExist(br *dto.CreateBranch) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := r.Branch.FindOne(ctx, bson.M{"branch_name": br.BranchName}).Decode(&br)
	if err != nil {
		return false
	}

	return true
}

func (r *DefaultRepo) OrganisationStatus(br *dto.OrganStatus) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.M{"_id": br.OrganisationID}
	update := bson.M{"$set": bson.M{"active": br.Active}}
	result, err := r.Organisation.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}
