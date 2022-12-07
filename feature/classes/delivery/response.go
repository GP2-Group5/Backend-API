package delivery

import (
	"github.com/GP2-Group5/Backend/feature/classes"
)

type ClassResp struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PIC          string `json:"pic"`
	Start_date   string `json:"start_date"`
	Graduated_at string `json:"graduated_at"`
}

func fromCore(dataCore classes.ClassCore) ClassResp {
	return ClassResp{
		ID:           int(dataCore.ID),
		Name:         dataCore.Name,
		PIC:          dataCore.Users,
		Start_date:   dataCore.Start_date,
		Graduated_at: dataCore.Graduated_at,
	}
}

func fromCoreList(dataCore []classes.ClassCore) []ClassResp {
	var dataResponse []ClassResp

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
