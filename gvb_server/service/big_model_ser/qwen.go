package big_model_ser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gvb_server/INIT/global"
	"gvb_server/models"
	"log"
	"net/http"
	"strings"
)

type QwenModel struct {
	SessionId uint `json:"session_id"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"` // 流式输出
}

type Response struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason interface{} `json:"finish_reason"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
	} `json:"choices"`
	Object            string      `json:"object"`
	Usage             interface{} `json:"usage"`
	Created           int         `json:"created"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
	Model             string      `json:"model"`
	Id                string      `json:"id"`
}

func (q QwenModel) Send(content string) (msgChan chan string, err error) {
	set := global.Config.BigModel.Setting
	msgChan = make(chan string)
	const baseURL = "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions"

	body := RequestBody{
		// 模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
		Model:    "qwen-plus",
		Messages: []Message{},
		Stream:   true,
	}

	// 查当前会话记录
	if q.SessionId != 0 {
		var sessionModel models.BigModelSessionModel
		var total int64
		global.DB.Preload("RoleModel").Take(&sessionModel, q.SessionId).Count(&total)
		if total == 0 {
			return nil, errors.New("不存在的会话")
		}
		body.Messages = append(body.Messages, Message{
			Role:    "system",
			Content: sessionModel.RoleModel.Prompt,
		})

		var chatList []models.BigModelChatModel
		global.DB.Order("created_at asc").Find(&chatList, "session_id = ?", q.SessionId)
		for _, model := range chatList {
			body.Messages = append(body.Messages,
				Message{Role: "user", Content: model.Content},
				Message{Role: "assistant", Content: model.BotContent},
			)
		}
	}
	body.Messages = append(body.Messages, Message{
		Role:    "user",
		Content: content,
	})

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 创建 POST 请求
	request, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
		return
	}
	// 设置请求头
	// 若没有配置环境变量，请用百炼API Key将下行替换为：apiKey := "sk-xxx"
	request.Header.Set("Authorization", "Bearer "+set.ApiKey)
	request.Header.Set("Content-Type", "application/json")
	// 发送请求
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}

	scan := bufio.NewScanner(response.Body) // 分片读
	scan.Split(bufio.ScanLines)             // 按行读

	go func() {
		for scan.Scan() {
			text := scan.Text()
			if text == "" {
				continue
			}
			if strings.HasPrefix(text, "data: [DONE]") {
				close(msgChan)
				return
			}

			var result Response
			err := json.Unmarshal([]byte(text[6:]), &result)
			if err != nil {
				fmt.Println("错误", text)
				log.Fatal(err)
				return
			}
			for _, choice := range result.Choices {
				t := choice.Delta.Content
				if t == "" {
					continue
				}
				msgChan <- t
			}
		}
	}()
	return
}
