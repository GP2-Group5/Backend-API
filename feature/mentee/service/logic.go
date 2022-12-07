package service

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/mentee"
	"github.com/go-playground/validator/v10"
)

type menteeService struct {
	menteeRepository mentee.IMenteeRepository
	validate         *validator.Validate
}

func New(repo mentee.IMenteeRepository) mentee.IMenteeService {
	return &menteeService{
		menteeRepository: repo,
		validate:         validator.New(),
	}
}

func (s *menteeService) Create(input mentee.MenteeCore) (err error) {
	_, errCreate := s.menteeRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (s *menteeService) GetAll() (data []mentee.MenteeCore, err error) {
	data, err = s.menteeRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *menteeService) GetByID(id int) (data mentee.MenteeCore, err error) {
	dataCore, errData := s.menteeRepository.GetByID(id)
	if errData != nil {
		return mentee.MenteeCore{}, errData
	}
	return dataCore, nil
}

func (s *menteeService) Update(data mentee.MenteeCore, id int) error {
	_, errUpdate := s.menteeRepository.Update(data, id)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (s *menteeService) Delete(dataCore mentee.MenteeCore, id int) error {
	_, errDelete := s.menteeRepository.Delete(dataCore, id)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
