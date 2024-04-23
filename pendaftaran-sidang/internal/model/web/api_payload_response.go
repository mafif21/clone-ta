package web

type UpdatePeminatanPayload struct {
	PeminatanId int `json:"peminatan_id"`
}

type StudentData struct {
	Nim    int `json:"nim"`
	UserId int `json:"user_id"`
	TeamId int `json:"team_id"`
}

type GetDetailStudentResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   StudentData `json:"data"`
}

type MemberTeamResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []struct {
		Nim    int `json:"nim"`
		UserId int `json:"user_id"`
		TeamId int `json:"team_id"`
		User   struct {
			Username string `json:"username"`
			Nama     string `json:"nama"`
		} `json:"user"`
	} `json:"data"`
}

type ResetTeamResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []int  `json:"data"`
}
