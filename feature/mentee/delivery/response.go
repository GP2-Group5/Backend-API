package delivery

import (
	"github.com/GP2-Group5/Backend/feature/mentee"
)

type MenteeResp struct {
	ID        int
	Name      string          `json:"name"`
	Address   string          `json:"address"`
	Email     string          `json:"email"`
	Gender    string          `json:"gender"`
	Phone     string          `json:"phone"`
	Type      string          `json:"type"`
	Major     string          `json:"major"`
	Graduate  int             `json:"graduate"`
	Status_id []mentee.Status `json:"status_id"`
	Class_id  []mentee.Class  `json:"class_id"`
}
