package delivery

import (
	"net/http"

	"github.com/GP2-Group5/Backend/feature/status"
	"github.com/GP2-Group5/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type StatusDelivery struct {
	statusService status.IStatusService
}

func New(service status.IStatusService, e *echo.Echo) {
	handler := &StatusDelivery{
		statusService: service,
	}
	e.GET("/menteestatus", handler.GetAll)
}

func (d *StatusDelivery) GetAll(c echo.Context) error {
	result, err := d.statusService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCoreList(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all data", dataResp))
}
