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
		err = fmt.Errorf("查询录取详情失败")
		return "", err
	}
	if len(acceptDetails) == 0 {
		err = fmt.Errorf("暂无" + major + "专业在" + province + "的录取信息")
		return "", err
	}
	question, err := MakeQuestion(acceptDetails, major, province, grade, rank)
	if err != nil {
		return "", fmt.Errorf("问题拼接有误")
	}
	reply, err := deepseekHandler.SendRequest(question)
	if err != nil {
		return "", fmt.Errorf("请求deepseek失败")
	}
	return reply, nil
}

func MakeQuestion(acceptDetails []models.AcceptDetail, major, province string, grade float64, rank int) (string, error) {
	question := "赣南师范大学近几年," + major + "专业在" + province + "的录取情况如下:\n"
	for _, detail := range acceptDetails {
		question += detail.Year + "年," + detail.Subject + detail.Batch + "录取人数：" + fmt.Sprintf("%d", detail.AcceptCount) + "人;\n"
		question += "最低分：" + fmt.Sprintf("%.2f", detail.LowestScore) + "分，最低名次：" + fmt.Sprintf("%d", detail.LowestRank) + "名;\n"
		question += "最高分：" + fmt.Sprintf("%.2f", detail.HighestScore) + "分，平均分：" + fmt.Sprintf("%.2f", detail.Average) + "分;\n"
		question += "控制线：" + detail.ControlLine + ";\n"
		question += "我今年的成绩是：" + fmt.Sprintf("%.2f", grade) + "分，在" + province + "的名次是：" + fmt.Sprintf("%d", rank) + ";\n"
	}
	question += "请根据以上数据,预测我录取赣南师范大学的概率。"
	return question, nil
}
