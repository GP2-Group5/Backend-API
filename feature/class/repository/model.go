package repository

import "gorm.io/gorm"

type ClassModel struct {
	gorm.Model
	Name         string
	User_id      Users
	Start_date   string
	Graduated_at string
}

type Users struct {
	ID   int
	Name string
}
