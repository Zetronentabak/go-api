package usecase

import (
	"context"
	"go-api/crud"
	"go-api/crud/entity"
	"log"

	"github.com/pkg/errors"
)

type CrudUseCase struct {
	config   *entity.EnvConfig
	crudRepo crud.CrudRepositoryI
}

func NewCrudUseCase(config *entity.EnvConfig, crudRepo crud.CrudRepositoryI) crud.CrudUseCaseI {
	return &CrudUseCase{config: config, crudRepo: crudRepo}
}
func (cuc *CrudUseCase) InsertDataUC(ctx context.Context, req entity.DataProductRequest) (resp bool, err error) { //check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}
	if req.ProductName == "" {
		err = errors.New("failed to add data product ")
		log.Println(err)
		return false, err
	} //insert data
	err = cuc.crudRepo.InsertData(ctx, req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cuc *CrudUseCase) GetDataUC(ctx context.Context) (resp entity.GetDataResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.crudRepo.GetAllData(ctx)
	if err != nil {
		log.Println("failed to show all data product with default log")
		return list, err
	}
	return list, err
}
func (cuc *CrudUseCase) GetADataUC(ctx context.Context, req entity.DataProductRequest) (resp entity.GetDataResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	data, err := cuc.crudRepo.GetSingleData(ctx, req)
	if err != nil {
		log.Println("failed to show single data product with default log")
		return data, err
	}
	return data, err
}

func (cuc *CrudUseCase) UpdateDataUC(ctx context.Context, req entity.DataProductRequest) (resp bool, err error) { //check if context is nil
	if ctx == nil {
		ctx = context.Background()
	} //update data
	err = cuc.crudRepo.UpdateData(ctx, req)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (cuc *CrudUseCase) DeleteDataUC(ctx context.Context, req entity.DataProductRequest) (resp bool, err error) { //check if context is nil
	if ctx == nil {
		ctx = context.Background()
	} //update data
	err = cuc.crudRepo.DeleteData(ctx, req)
	if err != nil {
		return false, err
	}
	return true, nil
}
