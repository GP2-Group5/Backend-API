package delivery

import (
	"net/http"

	"github.com/GP2-Group5/Backend/feature/auth"
	"github.com/GP2-Group5/Backend/utils/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService auth.ServiceInterface
}

func New(service auth.ServiceInterface, e *echo.Echo) {
	handler := &AuthHandler{
		authService: service,
	}
	e.POST("/auth", handler.Login)
}

func (handler *AuthHandler) Login(c echo.Context) error {
	reqBody := UserRequest{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed to bind data"))
	}

	Data, err := handler.authService.Login(reqBody.Email, reqBody.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed to get token data"+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login success", Data))
}
