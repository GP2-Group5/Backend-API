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
