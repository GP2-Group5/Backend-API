package delivery

import (
	"net/http"

	"github.com/GP2-Group5/Backend/feature/classes"
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
	e.POST("/class", handler.Create)
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
