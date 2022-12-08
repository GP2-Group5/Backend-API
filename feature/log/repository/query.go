package repository

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/log"
	"gorm.io/gorm"
)

type logRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) log.ILogRepository {
	return &logRepository{
		db: db,
	}
}

func (r *logRepository) Create(input log.LogCore) (row int, err error) {
	userGorm := LogCoreToModel(input)
	tx := r.db.Create(&userGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

func (r *logRepository) Update(data log.LogCore, id int) (row int, err error) {
	var log Logs
	gormLog := LogCoreToModel(data)

	tx := r.db.First(&log, id)

	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	tz := r.db.Model(&log).Updates(gormLog)
	if tz.Error != nil {
		return -1, errors.New("data not found")
	}

	if tz.RowsAffected == 0 {
		return 0, errors.New("error insert")
	}

	return int(tz.RowsAffected), nil
}

func (r *logRepository) Delete(data log.LogCore, id int) (row int, err error) {
	logGorm := LogCoreToModel(data)

	tx := r.db.Delete(&logGorm, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}

	return int(tx.RowsAffected), nil
}

func (r *logRepository) GetByID(id int) (data log.Mentee, err error) {
	var mentee Mentee

	tx := r.db.Model(&Mentee{}).Preload("Logs.Status").Find(&mentee)

	if tx.Error != nil {
		return data, tx.Error
	}

	dataCore := mentee.toCore()

	return dataCore, nil
}
