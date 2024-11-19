package big_model_ser

import (
	"errors"
	"gvb_server/INIT/global"
)

type BigModelInterface interface {
	Send(content string) (chan string, error)
}

func Send(sessionId uint, content string) (msgChan chan string, err error) {
	var ser BigModelInterface
	switch global.Config.BigModel.Setting.Name {
	case "qwen":
		ser = QwenModel{SessionId: sessionId}
	case "wenxin":
	case "xinghuo":
	case "tiangong":
	case "ChatGPT":
	default:
		return nil, errors.New("不支持的大模型")
	}
	return ser.Send(content)
}
