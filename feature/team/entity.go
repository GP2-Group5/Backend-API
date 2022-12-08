package team

type TeamCore struct {
	ID   uint
	Name string
}

type ITeamRepository interface {
	GetAll() (data []TeamCore, err error)
}

type ITeamService interface {
	GetAll() (data []TeamCore, err error)
}
