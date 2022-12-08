package delivery

import "github.com/GP2-Group5/Backend/feature/log"

type LogReq struct {
	Feedback  string `json:"feedback"`
	Status_id int    `json:"status_id"`
	Mentee_id int    `json:"mentee_id"`
}

func toCore(data LogReq) log.LogCore {
	return log.LogCore{
		Feedback: data.Feedback,
		StatusID: uint(data.Status_id),
		MenteeID: uint(data.Mentee_id),
	}
}
