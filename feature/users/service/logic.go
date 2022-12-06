package service

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/users"
	"github.com/GP2-Group5/Backend/utils/helper"
)

type userService struct {
	userRepository users.RepositoryInterface
}

func New(repo users.RepositoryInterface) users.ServiceInterface {
	return &userService{
		userRepository: repo,
	}
}

func (s *userService) Create(input users.UserCore) (err error) {
	if input.Full_Name == "" || input.Email == "" || input.Password == "" || input.Role == "" {
		return errors.New("semua field harus diisi")
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
