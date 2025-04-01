package consult

import (
	"fmt"
	"go_server/internal/handler/network/server"
	"go_server/internal/services"
	"strconv"
)

// 查询历年录取详情
func GetAcceptDetail(resp server.Response) {
	query := resp.Context.Request.URL.Query()
	major := query.Get("major")
	province := query.Get("province")
	year, err := strconv.Atoi(query.Get("year"))
	if err != nil {
		resp.Failed("param error")
	}
	details, err := services.GetAcceptDetailsList(major, province, year)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	resp.Res["list"] = details
	resp.Success("operate success")
}

// AI问答
func AskQuestion(resp server.Response) {
	param := struct {
		Question string `json:"question"`
		Model    string `json:"model"`
	}{}
	err := resp.Json(&param)
	if err != nil || param.Question == "" {
		resp.Failed("param error")
		return
	}
	reply, err := services.AskQuestion(param.Question, param.Model)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
	}
	resp.Res["reply"] = reply
	resp.Success("operate success")
}

// 录取预测
func PredictEnroll(resp server.Response) {
	param := struct {
		Major    string  `json:"major"`
		Province string  `json:"province"`
		Subject  string  `json:"subject"`
		Grade    float64 `json:"grade"`
		Rank     int     `json:"rank"`
	}{}
	err := resp.Json(&param)
	if err != nil || param.Major == "" || param.Province == "" || param.Subject == "" || param.Grade <= 0 || param.Rank <= 0 {
		resp.Failed("param error")
		return
	}
	reply, err := services.PredictEnroll(param.Major, param.Province, param.Subject, param.Grade, param.Rank)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	resp.Res["reply"] = reply
	resp.Success("operate success")
}
