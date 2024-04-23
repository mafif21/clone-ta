package helper

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"penjadwalan-sidang/internal/model/web"
	"sync"
)

const PengajuanURI = "https://4b4a-182-2-44-234.ngrok-free.app"

//const OldSofiURI = "https://sofi.my.id"

func ChangePengajuanStatus(pengajuanChannel chan<- *web.GetPengajuanDatas, feedback string, status string, workflowType string, memberId []int, token string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := resty.New()
	result := &web.GetPengajuanDatas{}

	_, _ = client.R().SetResult(result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(map[string]interface{}{
			"feedback":      feedback,
			"status":        status,
			"workflow_type": workflowType,
			"member_id":     memberId,
		}).
		Patch(PengajuanURI + "/api/pengajuan/change-status")

	pengajuanChannel <- result
}

func GetPengajuanUsers(memberId []int, token string) []int {
	client := resty.New()
	result := &web.GetIdUsersPengajuan{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(map[string]interface{}{
			"users_id": memberId,
		}).
		Post(PengajuanURI + "/api/pengajuan/users")

	if err != nil {
		fmt.Println(err)
	}
	return result.Data
}

func GetAllPengajuan(pengajuanChannel chan<- []int, token string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := resty.New()
	result := &web.GetPengajuanDatas{}

	_, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		Get(PengajuanURI + "/api/pengajuan/get")

	if err != nil {
		fmt.Println(err)
	}

	var pengajuansId []int
	if result.Data != nil {
		for _, data := range result.Data {
			pengajuansId = append(pengajuansId, data.PengajuanId)
		}
	} else {
		pengajuansId = append(pengajuansId, 0)
	}

	fmt.Println(pengajuansId)

	pengajuanChannel <- pengajuansId
}

func CreateNotification(memberId []int, title string, message string, url string, token string) {

	client := resty.New()
	result := &web.NotificationCreateResponse{}

	_, _ = client.R().SetResult(result).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(map[string]interface{}{
			"user_id": memberId,
			"title":   title,
			"message": message,
			"url":     url,
		}).
		Post(PengajuanURI + "/api/notification/create")
}
