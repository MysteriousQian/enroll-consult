package teacher

import (
	"go_server/internal/handler/network/server"
	"go_server/internal/services"
)

// 教师列表查询
func GetTeacherList(resp server.Response) {
	param := struct {
		Page     int    `json:"page"`
		PageSize int    `json:"page_size"`
		Name     string `json:"name"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	teacherList, err := services.GetTeacherList(param.Page, param.PageSize, param.Name)
	if err != nil {
		resp.Failed("获取失败")
		return
	}
	resp.Res["list"] = teacherList
	resp.Success("操作成功")
}

// 添加教师
func AddTeacher(resp server.Response) {
	param := struct {
		Name          string `json:"name"`
		Avatar        string `json:"avatar"`
		Subject       string `json:"subject"`
		Position      string `json:"position"`
		Qualification string `json:"qualification"`
		Honor         string `json:"honor"`
		Description   string `json:"description"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.AddTeacher(param.Name, param.Avatar, param.Subject, param.Position, param.Qualification, param.Honor, param.Description)
	if err != nil {
		resp.Failed("添加失败")
		return
	}
	resp.Success("操作成功")
}

// 编辑教师
func EditTeacher(resp server.Response) {
	param := struct {
		Id            int64  `json:"id"`
		Name          string `json:"name"`
		Avatar        string `json:"avatar"`
		Subject       string `json:"subject"`
		Position      string `json:"position"`
		Qualification string `json:"qualification"`
		Honor         string `json:"honor"`
		Description   string `json:"description"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.UpdateTeacher(param.Id, param.Name, param.Avatar, param.Subject, param.Position, param.Qualification, param.Honor, param.Description)
	if err != nil {
		resp.Failed("编辑失败")
		return
	}
	resp.Success("操作成功")
}

// 删除教师
func DeleteTeacher(resp server.Response) {
	param := struct {
		Id int64 `json:"id"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.DeleteTeacher(param.Id)
	if err != nil {
		resp.Failed("删除失败")
		return
	}
	resp.Success("操作成功")
}
