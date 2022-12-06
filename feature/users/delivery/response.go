package delivery

type UserRespon struct {
	ID        uint     `json:"id"`
	Full_Name string   `json:"full_name"`
	Email     string   `json:"email"`
	Status    string   `json:"status"`
	Role      string   `json:"role"`
	Team      TeamResp `json:"team"`
}

type TeamResp struct {
	ID   uint
	Name string
}
