package BY

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 计算md5
func (_Data) Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
