package delivery

type ClassReq struct {
	Name         string `json:"name" form:"name"`
	PIC          string `json:"pic" form:"pic"`
	Start_date   string `json:"start_date" form:"start_date"`
	Graduated_at string `json:"graduated_at" form:"graduated_at"`
}
