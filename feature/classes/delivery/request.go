package delivery

import "github.com/GP2-Group5/Backend/feature/classes"

type ClassReq struct {
	Name         string `json:"name" form:"name"`
	User_id      uint   `json:"user_id" form:"user_id"`
	Start_date   string `json:"start_date" form:"start_date"`
	Graduated_at string `json:"graduated_at" form:"graduated_at"`
}

func toCore(data ClassReq) classes.ClassCore {
	return classes.ClassCore{
		Name:         data.Name,
		User_id:      data.User_id,
		Start_date:   data.Start_date,
		Graduated_at: data.Graduated_at,
	}
}
