package service

import (
	"github.com/GP2-Group5/Backend/feature/status"
)

type statusService struct {
	statusRepository status.IStatusRepository
}

func New(repo status.IStatusRepository) status.IStatusService {
	return &statusService{
		statusRepository: repo,
	}
}

func (s *statusService) GetAll() (data []status.StatusCore, err error) {
	data, err = s.statusRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
