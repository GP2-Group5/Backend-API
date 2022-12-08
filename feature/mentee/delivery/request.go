package delivery

import "github.com/GP2-Group5/Backend/feature/mentee"

type MenteeReq struct {
	Name      string `json:"name" form:"name"`
	Address   string `json:"address" form:"address"`
	Email     string `json:"email" form:"email"`
	Gender    string `json:"gender" form:"gender"`
	Phone     string `json:"phone" form:"phone"`
	Type      string `json:"type" form:"type"`
	Major     string `json:"major" form:"major"`
	Graduate  uint   `json:"graduate" form:"graduate"`
	Status_id uint   `json:"status_id" form:"status_id"`
	Class_id  uint   `json:"class_id" form:"class_id"`
}

func toCore(data MenteeReq) mentee.MenteeCore {
	return mentee.MenteeCore{
		Name:      data.Name,
		Address:   data.Address,
		Email:     data.Email,
		Gender:    data.Gender,
		Phone:     data.Phone,
		Type:      data.Type,
		Major:     data.Major,
		Graduate:  data.Graduate,
		Status_id: data.Status_id,
		Class_id:  data.Class_id,
	}
}
