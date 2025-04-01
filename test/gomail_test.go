package test

import (
	"strings"
	"testing"
)

func TestGoMail(t *testing.T) {
	provinces := "安徽、澳门、北京、重庆、福建、甘肃、广东、广西、贵州、海南、河北、黑龙江、河南、湖北、湖南、江苏、江西、吉林、辽宁、内蒙古、宁夏、青海、陕西、山东、上海、山西、四川、台湾、天津、香港、新疆、西藏、云南、浙江"
	list := strings.Split(provinces, "、")
	for _, v := range list {
		println(v)
	}
}
