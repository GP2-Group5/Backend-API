package repository

import (
	"time"

	"github.com/GP2-Group5/Backend/feature/auth"
)

type Users struct {
	ID        uint
	FullName  string
	Password  string
	Email     string `gorm:"unique"`
	Status    string
	Role      string `gorm:"type:enum('user', 'admin');default:'default'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func ToCore(userModel Users) auth.Core {
	return auth.Core{
		ID:       userModel.ID,
		Role:     userModel.Role,
		Email:    userModel.Email,
		Password: userModel.Password,
	}
}

func ToCoreLogin(userModel Users) auth.Login {
	return auth.Login{
		ID:    userModel.ID,
		Role:  userModel.Role,
		Email: userModel.Email,
	}
}
