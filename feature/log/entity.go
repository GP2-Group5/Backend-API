package log

import "time"

type LogCore struct {
	ID        uint
	Feedback  string `validate:"required"`
	StatusID  uint   `validate:"required"`
	Status    string
	MenteeID  uint `validate:"required"`
	Mentee    string
	UsersID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status struct {
	ID   uint
	Name string
}

type Mentee struct {
	ID   uint
	Name string
	Log  []LogCore
}

type Users struct {
	ID   int
	Name string
}

type ILogRepository interface {
	Create(data LogCore) (row int, err error)
	Update(data LogCore, id int) (row int, err error)
	Delete(data LogCore, id int) (row int, err error)
	GetByID(id int) (data Mentee, err error)
}

type ILogService interface {
	Create(data LogCore) (err error)
	Update(data LogCore, id int) (err error)
	Delete(data LogCore, id int) (err error)
	GetByID(id int) (data Mentee, err error)
}
