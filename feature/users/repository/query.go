package repository

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/users"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(input users.UserCore) (row int, err error) {
	userGorm := UsersCoreToModel(input)
	tx := r.db.Create(&userGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}
