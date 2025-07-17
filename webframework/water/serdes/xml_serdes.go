package serdes

import "encoding/xml"

type XmlSerDes struct{}

var _ SerDes = (*XmlSerDes)(nil)

func NewXmlSerDes() *XmlSerDes {
	return &XmlSerDes{}
}

// Serialize implements SerDes.
func (x *XmlSerDes) Serialize(v any) ([]byte, error) {
	return xml.Marshal(v)
}

// Deserialize implements SerDes.
func (x *XmlSerDes) Deserialize(d []byte, v any) error {
	return xml.Unmarshal(d, v)
}
