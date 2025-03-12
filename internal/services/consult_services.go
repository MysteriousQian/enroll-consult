package services

import (
	"fmt"
	"go_server/internal/db/models"
	deepseekHandler "go_server/pkg/util/deepseek"
)

func AskQuestion(question string) (string, error) {
	reply, err := deepseekHandler.SendRequest(question)
	if err != nil {
		return "", err
	}
	return reply, nil
}

func PredictEnroll(major, province string, grade float64, rank int) (string, error) {
	acceptDetails, err := models.AcceptDetail{
		Major:    major,
		Province: province,
	}.SelectByMajorProvince()
	if err != nil {
		err = fmt.Errorf("failed to select accept detail")
		return "", err
	}
	if len(acceptDetails) == 0 {
		err = fmt.Errorf("no historical data available at the moment")
		return "", err
	}
	question, err := MakeQuestion(acceptDetails, major, province, grade, rank)
	reply, err := deepseekHandler.SendRequest(question)
	if err != nil {
		return "", err
	}
	return reply, nil
}

func MakeQuestion(acceptDetails []models.AcceptDetail, major, province string, grade float64, rank int) (string, error) {
	// question := "赣南师范大学近几年" + major + "专业在" + province + "的录取情况如下:\n"
	// for _, detail := range acceptDetails {

	// }
	return "", nil
}
