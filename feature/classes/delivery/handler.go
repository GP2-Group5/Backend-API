package delivery

import (
	"net/http"
	"strconv"

	"github.com/GP2-Group5/Backend/feature/classes"
	"github.com/GP2-Group5/Backend/middlewares"
	"github.com/GP2-Group5/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type ClassDelivery struct {
	classService classes.IClassService
}

func New(service classes.IClassService, e *echo.Echo) {
	handler := &ClassDelivery{
		classService: service,
	}
	e.POST("/class", handler.Create, middlewares.JWTMiddleware())
	e.GET("/class", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/class/:id", handler.GetByID, middlewares.JWTMiddleware())
	e.DELETE("/class/:id", handler.Delete, middlewares.JWTMiddleware())
	e.PUT("/class/:id", handler.Update, middlewares.JWTMiddleware())
}

func (d *ClassDelivery) GetAll(c echo.Context) error {
	result, err := d.classService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCoreList(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all data", dataResp))
}

func (d *ClassDelivery) Create(c echo.Context) error {
	classInput := ClassReq{}
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error bindingdata"+errBind.Error()))
	}

	dataCore := toCore(classInput)
	err := d.classService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success input class"))
}

func (d *ClassDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := d.classService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	dataResp := fromCore(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get class data by id", dataResp))
}

func (d *ClassDelivery) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := classes.ClassCore{}

	errResult := d.classService.Delete(user, id)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete class by id"))
}

func (d *ClassDelivery) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := classes.ClassCore{}

	c.Bind(&user)

	errResult := d.classService.Update(user, id)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update class by id"))
}
