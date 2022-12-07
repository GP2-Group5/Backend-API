package mentee

import "time"

type MenteeCore struct {
	ID         uint
	Name       string `validate:"required"`
	Address    string `validate:"required"`
	Email      string `validate:"required,email"`
	Gender     string `validate:"required"`
	Phone      string `validate:"required"`
	Type       string `validate:"required"`
	Major      string `validate:"required"`
	Graduate   uint   `validate:"required"`
	Status_id  uint   `validate:"required"`
	StatusName string
	Class_id   uint `validate:"required"`
	ClassName  string
	Deleted_at time.Time
	Created_at time.Time
	Updated_at time.Time
}

type Status struct {
	StatusID   int
	StatusName string
	Created_at time.Time
	Updated_at time.Time
}

type Class struct {
	ClassID   int
	ClassName string
}

type IMenteeRepository interface {
	Create(input MenteeCore) (row int, err error)
	GetAll() (data []MenteeCore, err error)
	GetByID(id int) (data MenteeCore, err error)
	Delete(data MenteeCore, id int) (row int, err error)
	Update(data MenteeCore, id int) (row int, err error)
}

type IMenteeService interface {
	Create(input MenteeCore) (err error)
	GetAll() (data []MenteeCore, err error)
	GetByID(id int) (data MenteeCore, err error)
	Delete(data MenteeCore, id int) (err error)
	Update(data MenteeCore, id int) (err error)
}
