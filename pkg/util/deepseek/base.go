package deepseekHandler

import (
	"context"
	"go_server/pkg/util/log"

	"github.com/p9966/go-deepseek"
	"github.com/spf13/viper"
)

/*
*
定义deepseek全局变量
*/
var DeepseekClient *deepseek.Client

/*
*
初始化deepseek客户端
*/
func Init() {
	apiKey := viper.GetString("deepseek.apiKey")
	if apiKey == "" {
		log.Error("deepseek apiKey not set")
		return
	}
	DeepseekClient = deepseek.NewClient("Bearer " + apiKey)
	log.Info("deepseek client initialized")
}

/*
*
向deepseek发送请求
*/
func SendRequest(question string) (string, error) {
	request := deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{
				Role:    deepseek.ChatMessageRoleSystem,
				Content: "你是一个强大的数据分析员",
			},
			{
				Role:    deepseek.ChatMessageRoleUser,
				Content: question,
			},
		},
	}
	ctx := context.Background()
	resp, err := DeepseekClient.CreateChatCompletion(ctx, &request)
	if err != nil {
		log.Error("ChatCompletion failed: %v", err)
		return "", err
	}
	if len(resp.Choices) == 0 {
		log.Error("No response choices available")
		return "", nil
	}
	return resp.Choices[0].Message.Content, nil
}
