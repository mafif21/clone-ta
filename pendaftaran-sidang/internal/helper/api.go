package helper

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"pendaftaran-sidang/internal/exception"
	"pendaftaran-sidang/internal/model/web"
	"strconv"
)

var apiUrl string = "https://sofi.my.id"

func GetDetailStudent(studentId int) (*web.GetDetailStudentResponse, *exception.ErrorResponse, error) {
	client := resty.New()
	result := &web.GetDetailStudentResponse{}

	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		Get(apiUrl + "/api/student/" + strconv.Itoa(studentId))

	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		errorResponse := &exception.ErrorResponse{}
		if err := json.Unmarshal(res.Body(), errorResponse); err != nil {
			return nil, nil, err
		}
		return nil, errorResponse, nil
	}

	return result, nil, nil
}

func UpdateUserTeamID(userID int, teamID int) (*web.GetDetailStudentResponse, *exception.ErrorResponse, error) {
	client := resty.New()
	result := &web.GetDetailStudentResponse{}

	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"team_id": teamID,
		}).
		Patch(apiUrl + "/api/student/team/update/" + strconv.Itoa(userID))

	if err != nil {
		return nil, nil, err
	}

	fmt.Println(res.StatusCode())
	if res.StatusCode() != 200 || res.IsError() {

		errorResponse := &exception.ErrorResponse{}
		if err := json.Unmarshal(res.Body(), errorResponse); err != nil {
			return nil, nil, err
		}
		return nil, errorResponse, nil
	}

	return result, nil, nil
}

func ResetTeam(teamID int) (*web.ResetTeamResponse, *exception.ErrorResponse, error) {
	client := resty.New()
	result := &web.ResetTeamResponse{}
	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"team_id": teamID,
		}).
		Patch(apiUrl + "/api/student/team/reset")

	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		errorResponse := &exception.ErrorResponse{}
		if err := json.Unmarshal(res.Body(), errorResponse); err != nil {
			return nil, nil, err
		}
		return nil, errorResponse, nil
	}

	return result, nil, nil
}

func GetAllTeamMember(teamId int) (*web.MemberTeamResponse, *exception.ErrorResponse, error) {
	client := resty.New()
	result := &web.MemberTeamResponse{}
	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		Get(apiUrl + "/api/student/team/" + strconv.Itoa(teamId))

	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		errorResponse := &exception.ErrorResponse{}
		if err := json.Unmarshal(res.Body(), errorResponse); err != nil {
			return nil, nil, err
		}
		return nil, errorResponse, nil
	}

	return result, nil, nil
}

func UpdatePeminatanByUserId(peminatanId int, userId int) (*web.GetDetailStudentResponse, *exception.ErrorResponse, error) {
	client := resty.New()
	result := &web.GetDetailStudentResponse{}

	res, err := client.R().SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"peminatan_id": peminatanId,
		}).Patch(apiUrl + "/api/sidang/update/" + strconv.Itoa(userId))

	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode() != 200 || res.IsError() {
		errorResponse := &exception.ErrorResponse{}
		if err := json.Unmarshal(res.Body(), errorResponse); err != nil {
			return nil, nil, err
		}
		return nil, errorResponse, nil
	}

	return result, nil, nil
}
