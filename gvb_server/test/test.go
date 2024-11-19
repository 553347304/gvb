package test

import (
	"fmt"
	"gvb_server/service/big_model_ser"
)

func Main() {
	send, err := big_model_ser.Send("qwen", "讲一个笑话")
	if err != nil {
		return
	}
	for msg := range send {
		fmt.Println(msg)
	}
}
