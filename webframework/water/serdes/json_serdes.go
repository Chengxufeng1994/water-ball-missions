package serdes

import "encoding/json"

type JsonSerDes struct{}

var _ SerDes = (*JsonSerDes)(nil)

func NewJsonSerDes() *JsonSerDes {
	return &JsonSerDes{}
}

func (j *JsonSerDes) Deserialize(d []byte, v any) error {
	return json.Unmarshal(d, v)
}

func (j *JsonSerDes) Serialize(v any) ([]byte, error) {
	return json.Marshal(v)
}
