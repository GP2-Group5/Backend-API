package repository

import (
	"github.com/GP2-Group5/Backend/feature/classes/repository"
	"github.com/GP2-Group5/Backend/feature/mentee"
	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name      string
	Address   string
	Email     string
	Gender    string
	Phone     string
	Type      string
	Major     string
	Graduate  uint
	Status_id uint
	Status    Status
	Class_id  uint
	Class     Class
	Log       []Log
}

type Status struct {
	gorm.Model
	Name string
}

type Class struct {
	gorm.Model
	Name         string
	User_id      int
	Start_date   string
	Graduated_at string
}

type Log struct {
	gorm.Model
	Feedback  string
	Status_id uint
	Status    Status
	Mentee_id uint
	Mentee    Mentee
	UsersID   uint
	Users     repository.Users
}

func MenteeCoreToModel(mentee mentee.MenteeCore) Mentee {
	menteeData := Mentee{
		Name:      mentee.Name,
		Address:   mentee.Address,
		Email:     mentee.Email,
		Gender:    mentee.Gender,
		Phone:     mentee.Phone,
		Type:      mentee.Type,
		Major:     mentee.Major,
		Graduate:  uint(mentee.Graduate),
		Status_id: uint(mentee.Status_id),
		Class_id:  uint(mentee.Class_id),
	}

	return menteeData
}

func (dataModel *Mentee) toCore() mentee.MenteeCore {
	return mentee.MenteeCore{
		ID:         dataModel.ID,
		Name:       dataModel.Name,
		Address:    dataModel.Address,
		Email:      dataModel.Email,
		Gender:     dataModel.Gender,
		Phone:      dataModel.Gender,
		Type:       dataModel.Type,
		Major:      dataModel.Major,
		Graduate:   dataModel.Graduate,
		Status_id:  dataModel.Status_id,
		Class_id:   dataModel.Class_id,
		StatusName: dataModel.Status.Name,
		ClassName:  dataModel.Class.Name,
		Log:        toCoreListLog(dataModel.Log),
	}
}

func toCoreList(dataModel []Mentee) []mentee.MenteeCore {
	var dataCore []mentee.MenteeCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

func (dataModel *Log) toCore() mentee.LogCore {
	return mentee.LogCore{
		ID:       dataModel.ID,
		Feedback: dataModel.Feedback,
		StatusID: dataModel.Status_id,
		Status:   dataModel.Status.Name,
		MenteeID: dataModel.Mentee_id,
		Mentee:   dataModel.Mentee.Name,
	}
}

func toCoreListLog(dataModel []Log) []mentee.LogCore {
	var dataCore []mentee.LogCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
