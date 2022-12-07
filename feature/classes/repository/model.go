package repository

import (
	"github.com/GP2-Group5/Backend/feature/classes"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name         string
	UsersID      uint
	Users        Users
	Start_date   string
	Graduated_at string
}

type Users struct {
	gorm.Model
	Email     string
	Full_Name string
	Password  string
	Status    string
	Role      string
	TeamID    uint
	Class     []Class
}

func ClassCoreToModel(class classes.ClassCore) Class {
	userData := Class{
		Name:         class.Name,
		UsersID:      class.User_id,
		Start_date:   class.Start_date,
		Graduated_at: class.Graduated_at,
	}

	return userData
}
