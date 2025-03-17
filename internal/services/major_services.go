package services

import (
	"fmt"
	"go_server/internal/db/models"
	"time"
)

func GetMajorList(page, size int, name string) (majorList []models.Major, err error) {
	majorList, err = models.Major{}.SelectMajors(page, size, name)
	if err != nil {
		err = fmt.Errorf("获取失败")
	}
	return
}

func AddMajor(name, employeDest, description string, popularity, workRate, studyRate float64) (err error) {
	now := time.Now().Unix()
	err = models.Major{
		Name:        name,
		EmployeDest: employeDest,
		Description: description,
		Popularity:  popularity,
		WorkRate:    workRate,
		StudyRate:   studyRate,
		CreateTime:  now,
		UpdateTime:  now,
	}.Create()
	if err != nil {
		err = fmt.Errorf("创建失败")
	}
	return
}

func EditMajor(id int64, name, employeDest, description string, popularity, workRate, studyRate float64) (err error) {
	now := time.Now().Unix()
	err = models.Major{
		Id:          id,
		Name:        name,
		EmployeDest: employeDest,
		Description: description,
		Popularity:  popularity,
		WorkRate:    workRate,
		StudyRate:   studyRate,
		UpdateTime:  now,
	}.Update()
	if err != nil {
		err = fmt.Errorf("编辑失败")
	}
	return
}

func DeleteMajor(id int64) (err error) {
	err = models.Major{
		Id: id,
	}.Delete()
	if err != nil {
		err = fmt.Errorf("删除失败")
	}
	return
}
