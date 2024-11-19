package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func SSEDomoView(c *gin.Context) {
	// 信道
	var channel = make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			time.Sleep(time.Second)
		}
		close(channel)
	}()
	// 流媒体
	c.Stream(func(w io.Writer) bool {
		s, ok := <-channel
		if ok {
			c.SSEvent("name", s)
			return true
		}
		return false
	})

}
func main() {
	r := gin.Default()
	r.GET("/sse", SSEDomoView)
	r.Run(":8081")
}
