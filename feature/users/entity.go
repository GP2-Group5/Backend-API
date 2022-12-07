package users

import "time"

type UserCore struct {
	ID        uint
	Full_Name string `validate:"required"`
	Password  string `validate:"required"`
	Email     string `validate:"required,email"`
	Status    string
	Role      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	TeamID    uint `validate:"required"`
	Team      string
}

type TeamCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type RepositoryInterface interface {
	GetAll() (data []UserCore, err error)
	GetByID(id int) (data UserCore, err error)
	Create(input UserCore) (row int, err error)
	Delete(data UserCore, id int) (row int, err error)
	Update(data UserCore, id int) (row int, err error)
}

type ServiceInterface interface {
	GetAll() (data []UserCore, err error)
	GetByID(id int) (data UserCore, err error)
	Create(input UserCore) (err error)
	Delete(data UserCore, id int) (err error)
	Update(data UserCore, id int) (err error)
}
