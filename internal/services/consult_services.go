package services

import (
	"fmt"
	"go_server/internal/db/models"
	deepseekHandler "go_server/pkg/util/deepseek"
	"go_server/pkg/util/log"
)

func GetAcceptDetailsList(major, province string, year int) (details []models.AcceptDetail, err error) {
	details, err = models.AcceptDetail{}.SelectAcceptDetailsList(major, province, fmt.Sprintf("%d", year))
	if err != nil {
		return []models.AcceptDetail{}, fmt.Errorf("查询录取详情失败")
	}
	return
}

func AskQuestion(question string) (string, error) {
	reply, err := deepseekHandler.SendRequest(question)
	if err != nil {
		return "", err
	}
	return reply, nil
}

func PredictEnroll(major, province, subject string, grade float64, rank int) (string, error) {
	acceptDetails, err := models.AcceptDetail{
		Major:    major,
		Province: province,
	}.SelectByMajorProvince()
	if err != nil {
		err = fmt.Errorf("查询录取详情失败")
		return "", err
	}
	if len(acceptDetails) == 0 {
		err = fmt.Errorf("暂无" + major + "专业在" + province + "的录取信息")
		return "", err
	}
	question, err := MakeQuestion(major, province, subject, grade, rank)
	if err != nil {
		return "", fmt.Errorf("问题拼接有误")
	}
	reply, err := deepseekHandler.SendRequest(question)
	if err != nil {
		return "", fmt.Errorf("请求deepseek失败")
	}
	log.Info("deepseek返回结果:" + reply)
	return reply, nil
}

func MakeQuestion(major, province, subject string, grade float64, rank int) (string, error) {
	question := fmt.Sprintf("我是一名%s的%s类高考考生,今年我的分数是:%f,排名是:%d,请快速根据近几年的数据计算我录取赣南师范大学%s专业的概率", province, subject, grade, rank, major)
	return question, nil
}
