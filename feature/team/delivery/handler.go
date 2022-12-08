package delivery

import (
	"net/http"

	"github.com/GP2-Group5/Backend/feature/team"
	"github.com/GP2-Group5/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type TeamDelivery struct {
	teamService team.ITeamService
}

func New(service team.ITeamService, e *echo.Echo) {
	handler := &TeamDelivery{
		teamService: service,
	}
	e.GET("/team", handler.GetAll)
}

func (d *TeamDelivery) GetAll(c echo.Context) error {
	result, err := d.teamService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCoreList(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all data", dataResp))
}
