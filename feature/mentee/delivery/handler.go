package delivery

import (
	"net/http"
	"strconv"

	"github.com/GP2-Group5/Backend/feature/mentee"
	"github.com/GP2-Group5/Backend/middlewares"
	"github.com/GP2-Group5/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeService mentee.IMenteeService
}

func New(service mentee.IMenteeService, e *echo.Echo) {
	handler := &MenteeDelivery{
		menteeService: service,
	}

	e.POST("/mentee", handler.Create, middlewares.JWTMiddleware())
	e.GET("/mentee", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/mentee/:id", handler.GetByID, middlewares.JWTMiddleware())
	e.DELETE("/mentee/:id", handler.Delete, middlewares.JWTMiddleware())
	e.PUT("/mentee/:id", handler.Update, middlewares.JWTMiddleware())
}

func (d *MenteeDelivery) Create(c echo.Context) error {
	menteeInput := MenteeReq{}
	errBind := c.Bind(&menteeInput)

	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Error binding data"+errBind.Error()))
	}

	dataCore := toCore(menteeInput)
	err := d.menteeService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success register mentee"))
}

func (d *MenteeDelivery) GetAll(c echo.Context) error {
	results, err := d.menteeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all data users", dataResponse))
}

func (d *MenteeDelivery) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := d.menteeService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCore(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get mentee data by id", dataResp))
}

func (d *MenteeDelivery) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	mentee := mentee.MenteeCore{}

	errResult := d.menteeService.Delete(mentee, id)
	if errResult != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete mentee by id"))
}

func (d *MenteeDelivery) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	mentee := mentee.MenteeCore{}

	c.Bind(&mentee)

	resultErr := d.menteeService.Update(mentee, id)

	if resultErr != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update mentee by ID"))
}
