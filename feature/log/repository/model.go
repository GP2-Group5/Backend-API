package repository

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Feedback  string
	Status_id Status
	Mentee_id Users
	User_id   int
	Proof     Proof
}

type Proof struct {
	LogID uint
	Link  string
}

type Status struct {
	ID   int
	Name string
}

type Mentee struct {
	ID   int
	Name string
}

type Users struct {
	ID   int
	Name string
}
