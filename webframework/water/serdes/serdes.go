package serdes

type (
	Serializer func(v any) ([]byte, error)

	Deserializer func(d []byte, v any) error

	SerDes interface {
		Serialize(v any) ([]byte, error)
		Deserialize(d []byte, v any) error
	}
)
