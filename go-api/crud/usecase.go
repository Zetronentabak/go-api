package crud

import (
	"context"
	"go-api/crud/entity"
)

type CrudUseCaseI interface {
	InsertDataUC(ctx context.Context, req entity.DataProductRequest) (resp bool, err error)
	GetDataUC(ctx context.Context) (resp entity.GetDataResponse, err error)
	GetADataUC(ctx context.Context, req entity.DataProductRequest) (resp entity.GetDataResponse, err error)
	UpdateDataUC(ctx context.Context, req entity.DataProductRequest) (resp bool, err error)
	DeleteDataUC(ctx context.Context, req entity.DataProductRequest) (resp bool, err error)
}
