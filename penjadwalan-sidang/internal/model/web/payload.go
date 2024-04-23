package web

type GetPengajuanData struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   struct {
		PengajuanId   int    `json:"id"`
		UserId        int    `json:"user_id"`
		Nim           string `json:"nim"`
		Pembimbing1Id int    `json:"pembimbing1_id"`
		Pembimbing2Id int    `json:"pembimbing2_id"`
		Status        string `json:"status"`
	} `json:"data"`
}

type GetPengajuanDatas struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []struct {
		PengajuanId   int    `json:"id"`
		UserId        int    `json:"user_id"`
		Nim           string `json:"nim"`
		Pembimbing1Id int    `json:"pembimbing1_id"`
		Pembimbing2Id int    `json:"pembimbing2_id"`
		Status        string `json:"status"`
	} `json:"data"`
}

type GetIdUsersPengajuan struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []int  `json:"data"`
}

type NotificationCreateResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []struct {
		Id int `json:"id"`
	} `json:"data"`
}

type MemberTeamResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []struct {
		UserId int `json:"user_id"`
	} `json:"data"`
}

type GetLectureByUserId struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   struct {
		UserId    int    `json:"user_id"`
		LectureId int    `json:"id"`
		Code      string `json:"code"`
		Jfa       string `json:"jfa"`
		Kk        string `json:"kk"`
	} `json:"data"`
}

type GetDetailStudent struct {
	Status  string `json:"status"`
	Message int    `json:"message"`
	Data    struct {
		Nim    int    `json:"nim"`
		UserId int    `json:"user_id"`
		TeamId int    `json:"team_id"`
		Kk     string `json:"kk"`
	} `json:"data"`
}
