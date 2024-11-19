package HTTP

import "time"

const timeout = 10 // 超时时间

type Response struct {
	URL    string            `json:"url"`    // URL
	Header map[string]string `json:"header"` // Header
	Data   any               `json:"data"`   // 传入数据
	Result any               `json:"result"` // 绑定结构体指针
	Time   time.Duration     `json:"time"`   // 超时时间
}

func (r *Response) Param() {
	// 默认超时时间
	if r.Time == 0 {
		r.Time = timeout
	}
}

type _Return struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}
