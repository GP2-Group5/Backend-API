package delivery

import (
	"net/http"
	"strconv"

	"github.com/GP2-Group5/Backend/feature/log"
	"github.com/GP2-Group5/Backend/middlewares"
	"github.com/GP2-Group5/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type LogDelivery struct {
	logService log.ILogService
}

func New(service log.ILogService, e *echo.Echo) {
	handler := &LogDelivery{
		logService: service,
	}
	e.POST("/log", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/log/:id", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/log/:id", handler.Delete, middlewares.JWTMiddleware())
	e.GET("/mentee/:id/log", handler.GetByID, middlewares.JWTMiddleware())
}

func (d *LogDelivery) Create(c echo.Context) error {
	logInput := LogReq{}
	errBind := c.Bind(&logInput)

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Error binding data"+errBind.Error()))
	}

	logInput.UsersID = middlewares.ExtractTokenUserId(c)

	dataCore := toCore(logInput)
	err := d.logService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success input log"))
}

func (d *LogDelivery) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log := LogReq{}

	errBind := c.Bind(&log)

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Error binding data"+errBind.Error()))
	}

	dataCore := toCore(log)

	resultErr := d.logService.Update(dataCore, id)

	if resultErr != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update log by ID"))
}

func (d *LogDelivery) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	log := log.LogCore{}

	errResult := d.logService.Delete(log, id)
	if errResult != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete log by id"))
}

func (d *LogDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := d.logService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCore(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get mentee data by id", dataResp))
}
