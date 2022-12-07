package delivery

import (
	"github.com/GP2-Group5/Backend/feature/users"
)

type UserRespon struct {
	ID        uint   `json:"id"`
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Role      string `json:"role"`
	Team      string `json:"team"`
}

func fromCore(dataCore users.UserCore) UserRespon {
	return UserRespon{
		ID:        dataCore.ID,
		Full_Name: dataCore.Full_Name,
		Email:     dataCore.Email,
		Status:    dataCore.Status,
		Role:      dataCore.Role,
		Team:      dataCore.Team,
	}
}

func fromCoreList(dataCore []users.UserCore) []UserRespon {
	var dataResponse []UserRespon

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
