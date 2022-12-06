package delivery

type MenteeReq struct {
	Name      string `json:"name" form:"name"`
	Address   string `json:"address" form:"address"`
	Email     string `json:"email" form:"email"`
	Gender    string `json:"gender" form:"gender"`
	Phone     string `json:"phone" form:"phone"`
	Type      string `json:"type" form:"type"`
	Major     string `json:"major" form:"major"`
	Graduate  int    `json:"graduate" form:"graduate"`
	Status_id int    `json:"status_id" form:"status_id"`
	Class_id  int    `json:"class_id" form:"class_id"`
}
