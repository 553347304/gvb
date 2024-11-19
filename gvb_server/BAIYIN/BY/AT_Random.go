package BY

import (
	"math/rand"
	"strconv"
)

// R 随机数字   范围|长度
func (_Random) R(num, length int) string {
	var ran string
	for i := 0; i < length; i++ {
		r := rand.Intn(num) // 随机数
		ran = ran + strconv.Itoa(r)
	}
	return ran
}

var animeName = []string{"白音", "小龙", "美雪", "天宇", "梦琪", "浩然", "雪莉"}

func (_Random) AnimeName() string {
	return animeName[rand.Intn(len(animeName))]
}
