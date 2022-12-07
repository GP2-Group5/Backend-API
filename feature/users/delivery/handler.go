package delivery

import (
	"strconv"

	"github.com/GP2-Group5/Backend/feature/users"
	"github.com/GP2-Group5/Backend/middlewares"
	"github.com/GP2-Group5/Backend/utils/helper"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService users.ServiceInterface
}

func New(service users.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}

	e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.POST("/users", handler.Create, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.Delete, middlewares.JWTMiddleware())
	e.PUT("/users/:id", handler.Update, middlewares.JWTMiddleware())
}

func (d *UserDelivery) GetAll(c echo.Context) error {
	result, err := d.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCoreList(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all data", dataResp))
}

func (d *UserDelivery) Create(c echo.Context) error {
	userInput := UserReq{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := d.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

func (d *UserDelivery) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := users.UserCore{}

	errResult := d.userService.Delete(user, id)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete user by id"))
}

func (d *UserDelivery) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := users.UserCore{}

	c.Bind(&user)

	errResult := d.userService.Update(user, id)
	if errResult != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("data not found"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update user by id"))
}
