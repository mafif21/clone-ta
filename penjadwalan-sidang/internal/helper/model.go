package helper

import (
	"penjadwalan-sidang/internal/model/entity"
	"penjadwalan-sidang/internal/model/web"
)

func ToScheduleResponse(schedule *entity.Schedule) web.ScheduleResponse {
	return web.ScheduleResponse{
		ID:               schedule.ID,
		PengajuanId:      schedule.PengajuanId,
		DateTime:         schedule.DateTime,
		Room:             schedule.Room,
		Penguji1Id:       schedule.Penguji1Id,
		Penguji2Id:       schedule.Penguji2Id,
		Status:           schedule.Status,
		Decision:         schedule.Decision,
		RevisionDuration: schedule.RevisionDuration,
		FlagAddRevision:  schedule.FlagAddRevision,
		FlagChangeScores: schedule.FlagChangeScores,
		CreatedAt:        schedule.CreatedAt,
		UpdatedAt:        schedule.UpdatedAt,
	}
}
