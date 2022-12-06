package service

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/auth"
	"github.com/GP2-Group5/Backend/middlewares"
	"github.com/GP2-Group5/Backend/utils/helper"
)

type authService struct {
	authData auth.RepositoryInterface
}

func New(data auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authData: data,
	}
}

func (service *authService) Login(email, password string) (LoginData auth.Login, err error) {

	if email == "" || password == "" {
		return LoginData, errors.New("email dan password harus diisi")
	}

	result, err := service.authData.FindUser(email, password)

	if err != nil {
		return LoginData, err
	}

	cekPass := helper.CheckPasswordHash(password, result.Password)

	if !cekPass {
		return LoginData, errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(result.ID), result.Role)
	if errToken != nil {
		return LoginData, errToken
	}

	var DataLogin auth.Login
	DataLogin.ID = result.ID
	DataLogin.Email = result.Email
	DataLogin.Role = result.Role
	DataLogin.Token = token

	return DataLogin, err
}
