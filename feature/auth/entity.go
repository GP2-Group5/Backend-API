package auth

import "time"

type Core struct {
	ID        uint
	FullName  string
	Password  string
	Email     string
	Status    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Login struct {
	ID    uint
	Email string
	Role  string
	Token string
}

type ServiceInterface interface {
	Login(email, password string) (LoginData Login, err error)
}

type RepositoryInterface interface {
	FindUser(email, password string) (loginData Core, err error)
}
