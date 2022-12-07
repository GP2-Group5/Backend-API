package service

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/classes"
	"github.com/go-playground/validator/v10"
)

type classService struct {
	classRepository classes.IClassRepository
	validate        *validator.Validate
}

func New(repo classes.IClassRepository) classes.IClassService {
	return &classService{
		classRepository: repo,
		validate:        validator.New(),
	}
}

func (s *classService) GetAll() (data []classes.ClassCore, err error) {
	data, err = s.classRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *classService) Create(input classes.ClassCore) (err error) {
	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	_, errCreate := s.classRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (s *classService) GetByID(id int) (data classes.ClassCore, err error) {
	dataCore, errData := s.classRepository.GetByID(id)
	if errData != nil {
		return classes.ClassCore{}, errData
	}
	return dataCore, nil
}

func (s *classService) Delete(data classes.ClassCore, id int) (err error) {
	_, errDelete := s.classRepository.Delete(data, id)
	if errDelete != nil {
		return errors.New("error delete")
	}

	return nil
}

func (s *classService) Update(data classes.ClassCore, id int) (err error) {
	_, errUpdate := s.classRepository.Update(data, id)
	if errUpdate != nil {
		return errors.New("error update")
	}

	return nil
}
