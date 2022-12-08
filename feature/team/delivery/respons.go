package delivery

import "github.com/GP2-Group5/Backend/feature/team"

type TeamResp struct {
	ID   uint
	Name string
}

func fromCore(dataCore team.TeamCore) TeamResp {
	return TeamResp{
		ID:   dataCore.ID,
		Name: dataCore.Name,
	}
}

func fromCoreList(dataCore []team.TeamCore) []TeamResp {
	var dataResponse []TeamResp

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
