package repository

import (
	"github.com/GP2-Group5/Backend/feature/status"
	"gorm.io/gorm"
)

type statusRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) status.IStatusRepository {
	return &statusRepository{
		db: db,
	}
}

func (r *statusRepository) GetAll() (data []status.StatusCore, err error) {
	var status []Status

	tx := r.db.Find(&status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dataCore := toCoreList(status)
	return dataCore, nil
}
