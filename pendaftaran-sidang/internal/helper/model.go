package helper

import (
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
)

func ToPengajuanResponse(pengajuan *entity.Pengajuan) web.PengajuanResponse {
	return web.PengajuanResponse{
		Id:             pengajuan.ID,
		UserId:         pengajuan.UserId,
		Nim:            pengajuan.Nim,
		Pembimbing1Id:  pengajuan.Pembimbing1Id,
		Pembimbing2Id:  pengajuan.Pembimbing2Id,
		Judul:          pengajuan.Judul,
		Eprt:           pengajuan.Eprt,
		DocTa:          pengajuan.DocTa,
		Makalah:        pengajuan.Makalah,
		Tak:            pengajuan.Tak,
		Status:         pengajuan.Status,
		SksLulus:       pengajuan.SksLulus,
		SksBelumLulus:  pengajuan.SksBelumLulus,
		IsEnglish:      pengajuan.IsEnglish,
		PeriodID:       pengajuan.PeriodID,
		SkPenguji:      pengajuan.SkPenguji,
		FormBimbingan1: pengajuan.FormBimbingan1,
		FormBimbingan2: pengajuan.FormBimbingan2,
		CreatedAt:      pengajuan.CreatedAt,
		Updated_at:     pengajuan.UpdatedAt,
	}
}

func ToPeriodResponse(period *entity.Period) web.PeriodResponse {
	if period.Pengajuans != nil {
		var pengajuans []web.PengajuanResponse
		for _, pengajuan := range period.Pengajuans {
			pengajuans = append(pengajuans, ToPengajuanResponse(&pengajuan))
		}

		return web.PeriodResponse{
			Id:          period.ID,
			Name:        period.Name,
			StartDate:   period.StartDate,
			EndDate:     period.EndDate,
			Description: period.Description,
			Pengajuans:  pengajuans,
			CreatedAt:   period.CreatedAt,
			UpdatedAt:   period.UpdatedAt,
		}
	} else {
		return web.PeriodResponse{
			Id:          period.ID,
			Name:        period.Name,
			StartDate:   period.StartDate,
			EndDate:     period.EndDate,
			Description: period.Description,
			CreatedAt:   period.CreatedAt,
			UpdatedAt:   period.UpdatedAt,
		}
	}

}

func ToTeamResponse(team *entity.Team) web.TeamResponse {
	return web.TeamResponse{
		Id:        team.ID,
		Name:      team.Name,
		CreatedAt: team.CreatedAt,
		UpdatedAt: team.UpdatedAt,
	}
}

func ToTeamResponseDetail(team *entity.Team, teamMembers *web.MemberTeamResponse) web.TeamResponseDetail {
	members := make([]web.MemberData, 0)
	for _, memberDataApi := range teamMembers.Data {
		member := web.MemberData{
			UserId:   memberDataApi.UserId,
			TeamId:   memberDataApi.TeamId,
			Nim:      memberDataApi.Nim,
			Username: memberDataApi.User.Username,
			Name:     memberDataApi.User.Nama,
		}

		members = append(members, member)
	}

	return web.TeamResponseDetail{
		Id:        team.ID,
		Name:      team.Name,
		Members:   members,
		CreatedAt: team.CreatedAt,
		UpdatedAt: team.UpdatedAt,
	}
}

func ToDocumentLogResponse(document *entity.DocumentLog) web.DocumentLogResponse {
	return web.DocumentLogResponse{
		Id:          document.ID,
		PengajuanId: document.PengajuanID,
		FileName:    document.FileName,
		Type:        document.Type,
		FileUrl:     document.FileUrl,
		CreatedBy:   document.CreatedBy,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
	}
}

func ToStatusLogResponse(statusLog *entity.StatusLog) web.StatusLogResponse {
	return web.StatusLogResponse{
		Id:           statusLog.ID,
		Feedback:     statusLog.Feedback,
		CreatedBy:    statusLog.CreatedBy,
		WorkFlowType: statusLog.WorkFlowType,
		Name:         statusLog.Name,
		PengajuanId:  statusLog.PengajuanID,
		CreatedAt:    statusLog.CreatedAt,
		UpdatedAt:    statusLog.UpdatedAt,
	}
}

func ToNotificationResponse(notification *entity.Notification) web.NotificationResponse {
	return web.NotificationResponse{
		Id:        notification.ID,
		UserId:    notification.UserId,
		Title:     notification.Title,
		Message:   notification.Message,
		Url:       notification.Url,
		ReadAt:    notification.ReadAt,
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}
