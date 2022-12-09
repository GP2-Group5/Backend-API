package delivery

import (
	"github.com/GP2-Group5/Backend/feature/mentee"
)

type MenteeResp struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Class   string `json:"class"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

type MenteeRespLog struct {
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Email   string    `json:"email"`
	Gender  string    `json:"gender"`
	Class   string    `json:"class"`
	Status  string    `json:"status"`
	Type    string    `json:"type"`
	Log     []LogResp `json:"log"`
}

type LogResp struct {
	ID       uint   `json:"id"`
	Feedback string `json:"feedback"`
	Status   string `json:"status"`
}

func fromCore(dataCore mentee.MenteeCore) MenteeResp {
	return MenteeResp{
		ID:      uint(dataCore.ID),
		Name:    dataCore.Name,
		Address: dataCore.Address,
		Email:   dataCore.Email,
		Gender:  dataCore.Gender,
		Type:    dataCore.Type,
		Status:  dataCore.StatusName,
		Class:   dataCore.ClassName,
	}
}

func fromCoreLg(dataCore mentee.MenteeCore) MenteeRespLog {
	return MenteeRespLog{
		ID:      uint(dataCore.ID),
		Name:    dataCore.Name,
		Address: dataCore.Address,
		Email:   dataCore.Email,
		Gender:  dataCore.Gender,
		Type:    dataCore.Type,
		Status:  dataCore.StatusName,
		Class:   dataCore.ClassName,
		Log:     fromCoreListLog(dataCore.Log),
	}
}

func fromCoreList(dataCore []mentee.MenteeCore) []MenteeResp {
	var dataResponse []MenteeResp

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreLog(dataCore mentee.LogCore) LogResp {
	return LogResp{
		ID:       uint(dataCore.ID),
		Feedback: dataCore.Feedback,
		Status:   dataCore.Status,
	}
}

func fromCoreListLog(dataCore []mentee.LogCore) []LogResp {
	var dataResponse []LogResp

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreLog(v))
	}
	return dataResponse
}
