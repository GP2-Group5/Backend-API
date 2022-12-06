package delivery

type LogReq struct {
	Feedback  string `json:"feedback"`
	User_id   int    `json:"user_id"`
	Status_id int    `json:"status_id"`
	Mentee_id int    `json:"mentee_id"`
	Proof     string `json:"proof"`
}
