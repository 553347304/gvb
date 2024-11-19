package FK

import (
	"gvb_server/BAIYIN/BY"
	"gvb_server/INIT/global"
)

func (_Redis) HGetAll(index string) map[string]int {
	var data = map[string]int{}
	list := global.Redis.HGetAll(index).Val() // 获取全部数据
	for id, val := range list {
		num := BY.Type.StringInt(val)
		data[id] = num
	}
	return data
}
func (_Redis) HGet(index string, id string) int {
	num, _ := global.Redis.HGet(index, id).Int()
	return num
}

func (_Redis) HSet(index string, id string, value interface{}) {
	global.Redis.HSet(index, id, value)
}

func (_Redis) Del(index string) {
	global.Redis.Del(index)
}

// HAdd 自增1
func (fk _Redis) HAdd(index string, id string, i int) {
	num := fk.HGet(index, id)
	fk.HSet(index, id, num+i)
}
