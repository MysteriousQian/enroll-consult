package services

import (
	"fmt"
	"go_server/internal/db/models"
	deepseekHandler "go_server/pkg/util/deepseek"
	"go_server/pkg/util/log"

	"github.com/p9966/go-deepseek"
)

// 获取历年录取数据
func GetAcceptDetailsList(major, province string, year int) (details []models.AcceptDetail, err error) {
	details, err = models.AcceptDetail{}.SelectAcceptDetailsList(major, province, fmt.Sprintf("%d", year))
	if err != nil {
		return []models.AcceptDetail{}, fmt.Errorf("查询录取详情失败")
	}
	return
}

// 问答功能
func AskQuestion(question, model string) (string, error) {
	if model == "" {
		model = deepseek.DeepSeekChat
	}
	if model != deepseek.DeepSeekChatR1 && model != deepseek.DeepSeekChat {
		return "", fmt.Errorf("模型类型错误")
	}
	reply, err := deepseekHandler.SendRequest(question, model)
	if err != nil {
		return "", err
	}
	return reply, nil
}

// 录取预测
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
	question, err := MakeQuestion(acceptDetails, major, province, subject, grade, rank)
	if err != nil {
		return "", fmt.Errorf("问题拼接有误")
	}
	log.Info("问题:" + question)
	reply, err := deepseekHandler.SendRequest(question, deepseek.DeepSeekChatR1)
	if err != nil {
		return "", fmt.Errorf("请求deepseek失败")
	}
	log.Info("deepseek返回结果:" + reply)
	return reply, nil
}

// 拼接问题
func MakeQuestion(acceptDetails []models.AcceptDetail, major, province, subject string, grade float64, rank int) (string, error) {
	temp := ""
	for _, detail := range acceptDetails {
		temp += fmt.Sprintf("%s,%s,%d,%d,%d;\n", detail.Year, detail.Subject, detail.LowestScore, detail.LowestRank, detail.HighestScore)
	}

	question := fmt.Sprintf(`我是一名%s的%s类高考考生,今年我的分数是:%d,排名是:%d,
		赣南师范大学近年,%s专业的各科类录取最低分、最低排名以及最高分,分别是:
		%s请根据近几年的数据详细分析并计算我录取赣南师范大学%s专业的概率,若概率较低推荐一下同校其他专业`, province, subject, int64(grade), rank, major,
		temp, major)
	return question, nil
}
