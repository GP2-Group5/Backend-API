package delivery

import (
	"github.com/GP2-Group5/Backend/feature/mentee"
)

type MenteeResp struct {
	ID       uint
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Type     string `json:"type"`
	Major    string `json:"major"`
	Graduate uint   `json:"graduate"`
	Status   string `json:"status"`
	Class    string `json:"class"`
}

func fromCore(dataCore mentee.MenteeCore) MenteeResp {
	return MenteeResp{
		ID:       uint(dataCore.ID),
		Name:     dataCore.Name,
		Address:  dataCore.Address,
		Email:    dataCore.Email,
		Gender:   dataCore.Gender,
		Phone:    dataCore.Phone,
		Type:     dataCore.Type,
		Major:    dataCore.Major,
		Graduate: uint(dataCore.Graduate),
		Status:   dataCore.StatusName,
		Class:    dataCore.ClassName,
	}
}

func fromCoreList(dataCore []mentee.MenteeCore) []MenteeResp {
	var dataResponse []MenteeResp

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
