package delivery

import "github.com/GP2-Group5/Backend/feature/log"

type LogReq struct {
	Feedback  string `json:"feedback"`
	Status_id uint   `json:"status_id"`
	Mentee_id uint   `json:"mentee_id"`
}

func toCore(data LogReq) log.LogCore {
	return log.LogCore{
		Feedback: data.Feedback,
		StatusID: data.Status_id,
		MenteeID: data.Mentee_id,
	}
}
