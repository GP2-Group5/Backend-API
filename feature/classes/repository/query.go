package repository

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/classes"
	"gorm.io/gorm"
)

type classRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) classes.IClassRepository {
	return &classRepository{
		db: db,
	}
}

func (r *classRepository) GetAll() (data []classes.ClassCore, err error) {
	var class []Class

	tx := r.db.Preload("Users").Find(&class)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dataCore := toCoreList(class)
	return dataCore, nil
}

func (r *classRepository) Create(input classes.ClassCore) (row int, err error) {
	classGorm := ClassCoreToModel(input)
	tx := r.db.Create(&classGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

func (r *classRepository) GetByID(id int) (data classes.ClassCore, err error) {
	var class Class

	tx := r.db.Preload("Users").First(&class, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	userCore := class.toCore()

	return userCore, nil
}

func (r *classRepository) Delete(data classes.ClassCore, id int) (row int, err error) {
	userGorm := ClassCoreToModel(data)

	tx := r.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}

	return int(tx.RowsAffected), nil
}

func (r *classRepository) Update(data classes.ClassCore, id int) (row int, err error) {
	var class Class
	gormUserCore := ClassCoreToModel(data)

	tx := r.db.First(&class, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	tz := r.db.Model(&class).Updates(gormUserCore)
	if tz.Error != nil {
		return -1, errors.New("data not found")
	}

	if tz.RowsAffected == 0 {
		return 0, err
	}

	return int(tz.RowsAffected), nil
}
