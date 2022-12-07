package repository

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/mentee"
	"gorm.io/gorm"
)

type menteeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.IMenteeRepository {
	return &menteeRepository{
		db: db,
	}
}

func (r *menteeRepository) Create(input mentee.MenteeCore) (row int, err error) {
	userGorm := MenteeCoreToModel(input)
	tx := r.db.Create(&userGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil

}

func (r *menteeRepository) GetAll() (data []mentee.MenteeCore, err error) {
	var mentees []Mentee

	tx := r.db.Preload("Status").Preload("Class").Find(&mentees)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(mentees)
	return dataCore, nil
}

func (r *menteeRepository) GetByID(id int) (data mentee.MenteeCore, err error) {
	var mentee Mentee

	tx := r.db.Preload("Status").Preload("Class").First(&mentee, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	dataCore := mentee.toCore()

	return dataCore, nil
}

func (r *menteeRepository) Delete(data mentee.MenteeCore, id int) (row int, err error) {
	userGorm := MenteeCoreToModel(data)

	tx := r.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}

	return int(tx.RowsAffected), nil
}

func (r *menteeRepository) Update(data mentee.MenteeCore, id int) (row int, err error) {
	var mentee Mentee
	gormUserCore := MenteeCoreToModel(data)

	tx := r.db.First(&mentee, id)

	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	tz := r.db.Model(&mentee).Updates(gormUserCore)
	if tz.Error != nil {
		return -1, errors.New("data not found")
	}

	if tz.RowsAffected == 0 {
		return 0, errors.New("error insert")
	}

	return int(tz.RowsAffected), nil
}
