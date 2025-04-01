package question

import (
	"fmt"
	"go_server/internal/handler/network/server"
	"go_server/internal/services"
	"strconv"
)

// 问题列表查询
func GetQuestionList(resp server.Response) {
	query := resp.Context.Request.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	pageSize, err := strconv.Atoi(query.Get("page_size"))
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	content := query.Get("content")
	questionList, total, err := services.GetQuestionList(page, pageSize, content)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	resp.Res["list"] = questionList
	resp.Res["total"] = total
	resp.Success("操作成功")
}

// 新增问题
func AddQuestion(resp server.Response) {
	param := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.AddQuestion(param.Title, param.Content)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	resp.Success("操作成功")
}

// 删除问题
func DeleteQuestion(resp server.Response) {
	param := struct {
		Id int64 `json:"id"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("参数错误")
		return
	}
	err = services.DeleteQuestion(param.Id)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	resp.Success("操作成功")
}
