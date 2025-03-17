package consult

import (
	"fmt"
	"go_server/internal/handler/network/server"
	"go_server/internal/services"
)

func GetAcceptDetail(resp server.Response) {
	param := struct {
		Major    string `json:"major"`
		Province string `json:"province"`
		Year     int    `json:"year"`
	}{}
	err := resp.Json(&param)
	if err != nil {
		resp.Failed("param error")
		return
	}
	details, err := services.GetAcceptDetailsList(param.Major, param.Province, param.Year)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	resp.Res["list"] = details
	resp.Success("operate success")
}

func AskQuestion(resp server.Response) {
	param := struct {
		Question string `json:"question"`
	}{}
	err := resp.Json(&param)
	if err != nil || param.Question == "" {
		resp.Failed("param error")
		return
	}
	reply, err := services.AskQuestion(param.Question)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
	}
	resp.Res["reply"] = reply
	resp.Success("operate success")
}

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
