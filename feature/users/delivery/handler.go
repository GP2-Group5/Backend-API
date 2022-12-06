package delivery

import (
	"github.com/GP2-Group5/Backend/feature/users"
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

	e.POST("/users", handler.Create)
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
