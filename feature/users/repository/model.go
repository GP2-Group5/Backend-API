package repository

import (
	"github.com/GP2-Group5/Backend/feature/classes/repository"
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
	Class     []repository.Class
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

func (dataModel *Users) toCore() users.UserCore {
	return users.UserCore{
		ID:        dataModel.ID,
		Full_Name: dataModel.Full_Name,
		Email:     dataModel.Email,
		Status:    dataModel.Status,
		Role:      dataModel.Role,
		Team:      dataModel.Team.Name,
	}
}

func toCoreList(dataModel []Users) []users.UserCore {
	var dataCore []users.UserCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
