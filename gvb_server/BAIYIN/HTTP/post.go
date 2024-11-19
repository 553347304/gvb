package HTTP

import (
	"gvb_server/BAIYIN/BY"
	"io"
	"net/http"
	"strings"
	"time"
)

func Post(R Response) _Return {
	R.Param()
	data := BY.Json.Marshal(R.Data)
	body := strings.NewReader(string(data))
	HTTP, err := http.NewRequest("POST", R.URL, body)
	if err != nil {
		return _Return{Code: 404, Body: err.Error()}
	}
	for key, value := range R.Header {
		HTTP.Header.Add(key, value)
	}
	// 超时时间
	client := http.Client{
		Timeout: time.Second * R.Time,
	}
	res, err := client.Do(HTTP)
	if err != nil {
		return _Return{Code: 404, Body: err.Error()}
	}
	// 转换JSON
	json, err := io.ReadAll(res.Body)
	if err != nil {
		return _Return{Code: 404, Body: err.Error()}
	}
	BY.Json.Unmarshal(json, R.Result)
	return _Return{
		Code: res.StatusCode,
		Body: string(json),
	}
}
