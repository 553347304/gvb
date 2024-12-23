package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1
	QiNiu ImageType = 2
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	switch s {
	case Local:
		return "本地"
	case QiNiu:
		return "七牛云"
	}
	return "其他"
}
