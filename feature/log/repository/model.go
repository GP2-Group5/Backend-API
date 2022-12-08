package repository

import (
	"github.com/GP2-Group5/Backend/feature/log"
	"gorm.io/gorm"
)

type Logs struct {
	gorm.Model
	Feedback string
	StatusID uint
	Status   Status
	MenteeID uint
	Mentee   Mentee
}

type Status struct {
	gorm.Model
	Name string
}

type Mentee struct {
	gorm.Model
	Name string
	Logs []Logs
}

func LogCoreToModel(data log.LogCore) Logs {
	menteeData := Logs{
		Feedback: data.Feedback,
		StatusID: data.StatusID,
		MenteeID: data.MenteeID,
	}

	return menteeData
}

func (dataModel *Mentee) toCore() log.Mentee {
	return log.Mentee{
		ID:   dataModel.ID,
		Name: dataModel.Name,
		Log:  toCoreList(dataModel.Logs),
	}
}

func (dataModel *Logs) toCoreLog() log.LogCore {
	return log.LogCore{
		ID:       dataModel.ID,
		Feedback: dataModel.Feedback,
		Status:   dataModel.Status.Name,
	}
}

func toCoreList(dataModel []Logs) []log.LogCore {
	var dataCore []log.LogCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCoreLog())
	}
	return dataCore
}
