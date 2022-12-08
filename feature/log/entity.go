package log

import "time"

type LogCore struct {
	ID        int
	Feedback  string `validate:"required"`
	StatusID  uint   `validate:"required"`
	MenteeID  uint   `validate:"required"`
	UsersID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status struct {
	ID   int
	Name string
}

type Mentee struct {
	ID   int
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
}

type ILogService interface {
	Create(data LogCore) (err error)
	Update(data LogCore, id int) (err error)
	Delete(data LogCore, id int) (err error)
}
