package repository

import (
	"github.com/GP2-Group5/Backend/feature/team"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name string
}

func (dataModel *Team) toCore() team.TeamCore {
	return team.TeamCore{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}

func toCoreList(dataModel []Team) []team.TeamCore {
	var dataCore []team.TeamCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
