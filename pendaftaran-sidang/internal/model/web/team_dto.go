package web

import "time"

type TeamCreateRequest struct {
	UserId int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type PersonalCreateRequest struct {
	UserId int    `validate:"required"`
	Name   string `validate:"required"`
}

type TeamUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `json:"name" validate:"required"`
}

type MemberData struct {
	UserId   int    `json:"user_id"`
	TeamId   int    `json:"team_id"`
	Nim      int    `json:"nim"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type TeamResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TeamResponseDetail struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Members   []MemberData `json:"members"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type AddMemberRequest struct {
	UserId int `json:"user_id" validate:"required"`
	TeamId int `json:"team_id" validate:"required"`
}

type LeaveMemberRequest struct {
	UserId int `json:"user_id" validate:"required"`
	TeamId int `json:"team_id" validate:"required"`
}

type AvailableMemberResponse struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
}
