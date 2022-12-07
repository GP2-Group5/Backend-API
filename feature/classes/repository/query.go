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
