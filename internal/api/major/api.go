package major

import (
	"go_server/internal/handler/network/server"
	"go_server/internal/services"
	"strconv"
)

// 获取专业列表
func GetMajorList(resp server.Response) {
	query := resp.Context.Request.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(query.Get("page_size"))
	if err != nil {
		pageSize = 10
	}
	name := query.Get("name")
	majorList, err := services.GetMajorList(page, pageSize, name)
	if err != nil {
		resp.Failed("获取失败")
		return
	}
	resp.Res["list"] = majorList
	resp.Success("操作成功")
}

// 添加专业
func AddMajor(resp server.Response) {
	param := struct {
		Name        string  `json:"name"`
		Popularity  float64 `json:"popularity"`
		EmployeDest string  `json:"employe_dest"`
		WorkRate    float64 `json:"work_rate"`
		StudyRate   float64 `json:"study_rate"`
		Description string  `json:"description"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.AddMajor(param.Name, param.EmployeDest, param.Description, param.Popularity, param.WorkRate, param.StudyRate)
	if err != nil {
		resp.Failed("添加失败")
		return
	}
	resp.Success("操作成功")
}

func EditMajor(resp server.Response) {
	param := struct {
		Id          int64   `json:"id"`
		Name        string  `json:"name"`
		Popularity  float64 `json:"popularity"`
		EmployeDest string  `json:"employe_dest"`
		WorkRate    float64 `json:"work_rate"`
		StudyRate   float64 `json:"study_rate"`
		Description string  `json:"description"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.EditMajor(param.Id, param.Name, param.EmployeDest, param.Description, param.Popularity, param.WorkRate, param.StudyRate)
	if err != nil {
		resp.Failed("编辑失败")
		return
	}
	resp.Success("操作成功")
}

func DeleteMajor(resp server.Response) {
	param := struct {
		Id int64 `json:"id"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.DeleteMajor(param.Id)
	if err != nil {
		resp.Failed("删除失败")
		return
	}
	resp.Success("操作成功")
}
