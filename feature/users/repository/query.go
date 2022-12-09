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

func (r *userRepository) GetAll() (data []users.UserCore, err error) {
	var users []Users

	tx := r.db.Preload("Team").Find(&users)
	if tx.Error != nil {
		return nil, errors.New("data not found")
	}
	dataCore := toCoreList(users)
	return dataCore, nil
}

func (r *userRepository) Create(input users.UserCore) (row int, err error) {
	userGorm := UsersCoreToModel(input)
	tx := r.db.Create(&userGorm)
	if tx.Error != nil {
		return -1, errors.New("insert failed")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

func (r *userRepository) GetByID(id int) (data users.UserCore, err error) {
	var user Users

	tx := r.db.Preload("Team").First(&user, id)

	if tx.Error != nil {
		return data, errors.New("data not found")
	}

	userCore := user.toCore()

	return userCore, nil
}

func (r *userRepository) Delete(data users.UserCore, id int) (row int, err error) {
	userGorm := UsersCoreToModel(data)

	tx := r.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}

	return int(tx.RowsAffected), nil
}

func (r *userRepository) Update(data users.UserCore, id int) (row int, err error) {
	var user Users
	gormUserCore := UsersCoreToModel(data)

	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	tz := r.db.Model(&user).Updates(gormUserCore)
	if tz.Error != nil {
		return -1, errors.New("data not found")
	}

	if tz.RowsAffected == 0 {
		return 0, err
	}

	return int(tz.RowsAffected), nil
}
