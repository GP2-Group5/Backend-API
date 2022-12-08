package service

import (
	"errors"

	"github.com/GP2-Group5/Backend/feature/log"
	"github.com/go-playground/validator"
)

type logService struct {
	logRepository log.ILogRepository
	validate      *validator.Validate
}

func New(repo log.ILogRepository) log.ILogService {
	return &logService{
		logRepository: repo,
		validate:      validator.New(),
	}
}

func (s *logService) Create(input log.LogCore) (err error) {
	if errValidate := s.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	_, errCreate := s.logRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (s *logService) Update(data log.LogCore, id int) error {
	_, errUpdate := s.logRepository.Update(data, id)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (s *logService) Delete(dataCore log.LogCore, id int) error {
	_, errDelete := s.logRepository.Delete(dataCore, id)
	if errDelete != nil {
		return errDelete
	}

	return nil
}
