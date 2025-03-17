package services

import (
	"fmt"
	"go_server/internal/db/models"
	"time"
)

func GetTeacherList(page, size int, name string) (teacherList []models.Teacher, err error) {
	teacherList, err = models.Teacher{}.SelectTeachers(page, size, name)
	return
}

func AddTeacher(name, avatar, subject, qualification, description string) (err error) {
	now := time.Now().Unix()
	err = models.Teacher{
		Name:          name,
		Avatar:        avatar,
		Subject:       subject,
		Qualification: qualification,
		Description:   description,
		CreateTime:    now,
		UpdateTime:    now,
	}.Create()
	if err != nil {
		err = fmt.Errorf("添加失败")
	}
	return
}

func UpdateTeacher(id int64, name, avatar, subject, qualification, description string) (err error) {
	now := time.Now().Unix()
	err = models.Teacher{
		Id:            id,
		Name:          name,
		Avatar:        avatar,
		Subject:       subject,
		Qualification: qualification,
		Description:   description,
		UpdateTime:    now,
	}.Update()
	if err != nil {
		err = fmt.Errorf("修改失败")
	}
	return
}

func DeleteTeacher(id int64) (err error) {
	err = models.Teacher{
		Id: id,
	}.Delete()
	if err != nil {
		err = fmt.Errorf("删除失败")
	}
	return
}
