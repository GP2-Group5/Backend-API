package service

import "github.com/GP2-Group5/Backend/feature/team"

type teamService struct {
	teamRepository team.ITeamRepository
}

func New(repo team.ITeamRepository) team.ITeamService {
	return &teamService{
		teamRepository: repo,
	}
}

func (s *teamService) GetAll() (data []team.TeamCore, err error) {
	data, err = s.teamRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
