package repository

import (
	"context"
	"fmt"
	"go-api/crud"
	"go-api/crud/entity"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CrudRepository struct{ mongoDB *mongo.Database }

func NewCrudRepository(mongo *mongo.Database) crud.CrudRepositoryI {
	return &CrudRepository{mongoDB: mongo}
}
func (cr CrudRepository) InsertData(ctx context.Context, req entity.DataProductRequest) (err error) {
	dataReq := bson.M{"product_name": req.ProductName}
	query, err := cr.mongoDB.Collection("product").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}
	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		productID := oid.Hex()
		dataUpdateProductID := bson.M{"_id": oid}
		dataObjectID := bson.M{"$set": bson.M{"product_id": productID}}
		_, err := cr.mongoDB.Collection("product").UpdateOne(ctx, dataUpdateProductID, dataObjectID)
		if err != nil {
			log.Println("error")
		}
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID ", err))
		log.Println("error")
	}
	return err
}

func (cr CrudRepository) GetAllData(ctx context.Context) (crudResp entity.GetDataResponse, err error) {
	query, err := cr.mongoDB.Collection("product").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return entity.GetDataResponse{}, err
	}
	defer query.Close(ctx)
	listDataProduct := make([]entity.DataProduct, 0)
	for query.Next(ctx) {
		var row entity.DataProduct
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataProduct = append(listDataProduct, row)
	}
	crudResp = entity.GetDataResponse{Data: listDataProduct}
	return crudResp, err
}
func (cr CrudRepository) GetSingleData(ctx context.Context, req entity.DataProductRequest) (crudResp entity.GetDataResponse, err error) {
	dataGetProductID := bson.M{"product_id": req.ProductID}
	query := cr.mongoDB.Collection("product").FindOne(ctx, dataGetProductID)
	listDataProduct := make([]entity.DataProduct, 0)
	var row entity.DataProduct
	err = query.Decode(&row)
	if err != nil {
		log.Println("error")
	}
	listDataProduct = append(listDataProduct, row)
	crudResp = entity.GetDataResponse{Data: listDataProduct}
	return crudResp, err
}

func (cr CrudRepository) UpdateData(ctx context.Context, req entity.DataProductRequest) (err error) {
	dataUpdateProductID := bson.M{"product_id": req.ProductID}
	dataObjectID := bson.M{"$set": bson.M{"product_name": req.ProductName}}
	_, err = cr.mongoDB.Collection("product").UpdateOne(ctx, dataUpdateProductID, dataObjectID)
	if err != nil {
		log.Println("error")
	}
	return err
}
func (cr CrudRepository) DeleteData(ctx context.Context, req entity.DataProductRequest) (err error) {
	dataUpdateProductID := bson.M{"product_id": req.ProductID}
	_, err = cr.mongoDB.Collection("product").DeleteOne(ctx, dataUpdateProductID)
	if err != nil {
		log.Println("error")
	}
	return err
}
