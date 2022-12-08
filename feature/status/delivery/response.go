package delivery

import (
	"github.com/GP2-Group5/Backend/feature/status"
)

type StatusResp struct {
	ID   uint
	Name string
}

func fromCore(dataCore status.StatusCore) StatusResp {
	return StatusResp{
		ID:   dataCore.ID,
		Name: dataCore.Name,
	}
}

func fromCoreList(dataCore []status.StatusCore) []StatusResp {
	var dataResponse []StatusResp

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
