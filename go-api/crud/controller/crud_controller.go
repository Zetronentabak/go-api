package controller

import (
	"context"
	"encoding/json"
	"go-api/crud"
	"go-api/crud/entity"

	"github.com/labstack/echo/v4"
)

type CrudController struct {
	e       *echo.Echo
	usecase crud.CrudUseCaseI
}

func NewCrudController(e *echo.Echo, usecase crud.CrudUseCaseI) *CrudController {
	return &CrudController{e: e, usecase: usecase}
}
func (cc *CrudController) InsertData(ec echo.Context) error {
	var req entity.DataProductRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	_, err = cc.usecase.InsertDataUC(context.Background(), req)
	if err != nil {
		return err
	}
	return ec.JSON(200, map[string]string{"message": "insert data successfully"})
}

func (cc *CrudController) GetData(ec echo.Context) error {
	data, err := cc.usecase.GetDataUC(context.Background())
	if err != nil {
		return err
	}
	return ec.JSON(200, data)
}
func (cc *CrudController) GetAData(ec echo.Context) error {
	var req entity.DataProductRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	data, err := cc.usecase.GetADataUC(context.Background(), req)
	if err != nil {
		return err
	}
	return ec.JSON(200, data)
}
func (cc *CrudController) UpdateData(ec echo.Context) error {
	var req entity.DataProductRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	_, err = cc.usecase.UpdateDataUC(context.Background(), req)
	if err != nil {
		return err
	}
	return ec.JSON(200, map[string]string{"message": "update data successfully"})
}
func (cc *CrudController) DeleteData(ec echo.Context) error {
	var req entity.DataProductRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	_, err = cc.usecase.DeleteDataUC(context.Background(), req)
	if err != nil {
		return err
	}
	return ec.JSON(200, map[string]string{"message": "delete data successfully"})
}
