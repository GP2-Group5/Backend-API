package delivery

import (
	"time"

	"github.com/GP2-Group5/Backend/feature/users"
)

type UserRespon struct {
	ID        uint      `json:"id"`
	Full_Name string    `json:"full_name"`
	Email     string    `json:"email"`
	Team      string    `json:"team"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func fromCore(dataCore users.UserCore) UserRespon {
	return UserRespon{
		ID:        dataCore.ID,
		Full_Name: dataCore.Full_Name,
		Email:     dataCore.Email,
		Team:      dataCore.Team,
		Status:    dataCore.Status,
		Role:      dataCore.Role,
		CreatedAt: dataCore.CreatedAt,
	}
}

func fromCoreList(dataCore []users.UserCore) []UserRespon {
	var dataResponse []UserRespon

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
