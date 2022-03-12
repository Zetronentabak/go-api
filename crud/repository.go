package crud

import (
	"context"
	"go-api/crud/entity"
)

type CrudRepositoryI interface {
	InsertData(ctx context.Context, req entity.DataProductRequest) error
	GetAllData(ctx context.Context) (crudResp entity.GetDataResponse, err error)
	GetSingleData(ctx context.Context, req entity.DataProductRequest) (crudResp entity.GetDataResponse, err error)
	UpdateData(ctx context.Context, req entity.DataProductRequest) error
	DeleteData(ctx context.Context, req entity.DataProductRequest) error
}
