package repository

import (
	"github.com/GP2-Group5/Backend/feature/auth"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) FindUser(email, password string) (loginData auth.Core, err error) {

	userModel := Users{}

	tx := repo.db.Where("email= ? and deleted_at is null", email).First(&userModel)

	err = tx.Error

	if err != nil {
		return loginData, err
	}

	loginData = ToCore(userModel)

	return loginData, nil

	// var userData Users
	// tx := repo.db.Where("email = ?", email).First(&userData)
	// if tx.Error != nil {
	// 	return "", tx.Error
	// }

	// if tx.RowsAffected == 0 {
	// 	return "", errors.New("login failed")
	// }

	// cekpass := helper.CheckPasswordHash(password, userData.Password)
	// if !cekpass {
	// 	return "", errors.New("login failed")
	// }

	// token, errToken := middlewares.CreateToken(int(userData.ID), userData.Role)
	// if errToken != nil {
	// 	return "", errToken
	// }

	// return token, nil
}
