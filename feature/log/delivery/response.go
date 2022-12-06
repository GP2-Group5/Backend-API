package delivery

import (
	"github.com/GP2-Group5/Backend/feature/log"
)

type LogResp struct {
	Mentee    []log.User
	ID        int    `json:"id"`
	Feedback  int    `json:"feedback"`
	Status_id int    `json:"status_id"`
	Proof     string `json:"proof"`
}
