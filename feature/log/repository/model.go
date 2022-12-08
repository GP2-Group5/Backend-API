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
}

func LogCoreToModel(data log.LogCore) Logs {
	menteeData := Logs{
		Feedback: data.Feedback,
		StatusID: data.StatusID,
		MenteeID: data.MenteeID,
	}

	return menteeData
}
