package repository

import (
	"github.com/GP2-Group5/Backend/feature/status"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Name string
}

func (dataModel *Status) toCore() status.StatusCore {
	return status.StatusCore{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}

func toCoreList(dataModel []Status) []status.StatusCore {
	var dataCore []status.StatusCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
