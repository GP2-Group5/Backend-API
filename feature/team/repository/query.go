package repository

import (
	"github.com/GP2-Group5/Backend/feature/team"
	"gorm.io/gorm"
)

type teamRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) team.ITeamRepository {
	return &teamRepository{
		db: db,
	}
}

func (r *teamRepository) GetAll() (data []team.TeamCore, err error) {
	var team []Team

	tx := r.db.Find(&team)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dataCore := toCoreList(team)
	return dataCore, nil
}
