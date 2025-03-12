package consult

import (
	"fmt"
	"go_server/internal/handler/network/server"
	"go_server/internal/services"
)

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
		Grade    float64 `json:"grade"`
		Rank     int     `json:"rank"`
	}{}
	err := resp.Json(&param)
	if err != nil || param.Major == "" || param.Province == "" || param.Grade <= 0 || param.Rank <= 0 {
		resp.Failed("param error")
		return
	}
	reply, err := services.PredictEnroll(param.Major, param.Province, param.Grade, param.Rank)
	if err != nil {
		resp.Failed(fmt.Sprintf("%v", err))
		return
	}
	resp.Res["reply"] = reply
	resp.Success("operate success")
}
