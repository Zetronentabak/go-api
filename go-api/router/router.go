package router

import (
	"go-api/crud"
	"go-api/crud/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, crudUseCase crud.CrudUseCaseI) {
	crudCtrl := controller.NewCrudController(e, crudUseCase)
	r := e.Group("/api")
	r.POST("/add", crudCtrl.InsertData)
	r.GET("/list", crudCtrl.GetData)
	r.GET("/get", crudCtrl.GetAData)
	r.PUT("/edit", crudCtrl.UpdateData)
	r.DELETE("/delete", crudCtrl.DeleteData)
}
