package delivery

type LogResp struct {
	Name      string `json:"name"`
	ID        uint   `json:"id"`
	Feedback  int    `json:"feedback"`
	Status_id int    `json:"status_id"`
}
