package repository

import (
	"github.com/GP2-Group5/Backend/feature/users"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email     string
	Full_Name string
	Password  string
	Status    string
	Role      string
	TeamID    uint
	Team      Team
}

type Team struct {
	gorm.Model
	Name string
}

func UsersCoreToModel(user users.UserCore) Users {
	userData := Users{
		Full_Name: user.Full_Name,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		Role:      user.Role,
		TeamID:    user.TeamID,
	}

	return userData
}
