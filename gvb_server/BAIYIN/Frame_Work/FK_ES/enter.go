package FK_ES

type TypeResponseList[T any] struct {
	Count   int64  `json:"total"` // 总数
	List    []T    `json:"list"`  // 值
	Message string `json:"message"`
}

const ID = "_id"
const KEY = "key"
