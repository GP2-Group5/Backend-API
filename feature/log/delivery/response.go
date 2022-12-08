package delivery

import "github.com/GP2-Group5/Backend/feature/log"

type MenteeResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Log  []LogResp
}

type LogResp struct {
	ID       uint   `json:"id"`
	Status   string `json:"status"`
	Feedback string `json:"feedback"`
}

func fromCore(dataCore log.Mentee) MenteeResp {
	return MenteeResp{
		ID:   dataCore.ID,
		Name: dataCore.Name,
		Log:  fromCoreListLog(dataCore.Log),
	}
}

func fromCoreLog(dataCore log.LogCore) LogResp {
	return LogResp{
		ID:       dataCore.ID,
		Status:   dataCore.Status,
		Feedback: dataCore.Feedback,
	}
}

func fromCoreListLog(dataCore []log.LogCore) []LogResp {
	var dataResponse []LogResp

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreLog(v))
	}
	return dataResponse
}
