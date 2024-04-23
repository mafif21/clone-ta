package services

import "pendaftaran-sidang/internal/model/web"

type TeamService interface {
	Create(request *web.TeamCreateRequest) (*web.TeamResponse, error)
	AddMember(request *web.AddMemberRequest) error
	LeaveTeam(request *web.LeaveMemberRequest) error
	Update(request *web.TeamUpdateRequest) (*web.TeamResponse, error)
	Delete(teamId int, userId int) error
	GetTeamByUserId(teamId int) (*web.TeamResponseDetail, error)
}
