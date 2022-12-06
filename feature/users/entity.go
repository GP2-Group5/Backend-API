package users

import "time"

type UserCore struct {
	ID        uint
	Full_Name string
	Password  string
	Email     string
	Status    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	TeamID    uint
	Team      TeamCore
}

type TeamCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type RepositoryInterface interface {
	Create(input UserCore) (row int, err error)
}

type ServiceInterface interface {
	Create(input UserCore) (err error)
}
