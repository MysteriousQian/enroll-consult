package services

import (
	"fmt"
	"go_server/internal/db/models"
	"time"
)

// 获取问题列表
func GetQuestionList(page, size int, content string) (questionList []models.Question, total int64, err error) {
	questionList, total, err = models.Question{}.SelectAllByPage(page, size, content)
	if err != nil {
		err = fmt.Errorf("获取失败")
	}
	return
}

// 添加问题
func AddQuestion(title, content string) (err error) {
	now := time.Now().Unix()
	err = models.Question{
		Title:      title,
		Content:    content,
		CreateTime: now,
		UpdateTime: now,
	}.Create()
	if err != nil {
		err = fmt.Errorf("添加失败")
	}
	return
}

// 删除问题
func DeleteQuestion(id int64) (err error) {
	err = models.Question{Id: id}.Delete()
	if err != nil {
		err = fmt.Errorf("删除失败")
	}
	return
}
