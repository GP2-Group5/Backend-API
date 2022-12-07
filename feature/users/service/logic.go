package service

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/users"
	"github.com/GP2-Group5/Backend/utils/helper"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository users.RepositoryInterface
	validate       *validator.Validate
}

func New(repo users.RepositoryInterface) users.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

func (s *userService) GetAll() (data []users.UserCore, err error) {
	data, err = s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *userService) GetByID(id int) (data users.UserCore, err error) {
	dataCore, errData := s.userRepository.GetByID(id)
	if errData != nil {
		return users.UserCore{}, errData
	}
	return dataCore, nil
}

func (s *userService) Create(input users.UserCore) (err error) {
	// if input.Full_Name == "" || input.Email == "" || input.Password == "" || input.Role == "" {
	// 	return errors.New("semua field harus diisi")
	// }
	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	input.Status = "Active"

	input.Password, err = helper.HashPassword(input.Password)
	if err != nil {
		return errors.New("failed to hash")
	}

	_, errCreate := s.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (s *userService) Delete(data users.UserCore, id int) (err error) {
	_, errDelete := s.userRepository.Delete(data, id)
	if errDelete != nil {
		return errors.New("error delete")
	}

	return nil
}

func (s *userService) Update(data users.UserCore, id int) (err error) {
	_, errUpdate := s.userRepository.Update(data, id)
	if errUpdate != nil {
		return errors.New("error update")
	}

	return nil
}
