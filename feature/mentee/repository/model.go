package repository

import "time"

type MenteeModel struct {
	Name       string
	Address    string
	Email      string
	Gender     string
	Phone      string
	Type       string
	Major      string
	Graduate   int
	Status_id  []Status
	Class_id   []Class
	Deleted_at time.Time
	Created_at time.Time
	Updated_at time.Time
}

type Status struct {
	ID         int
	Name       int
	Created_at time.Time
	Updated_at time.Time
}

type Class struct {
	ID   int
	Name string
}
